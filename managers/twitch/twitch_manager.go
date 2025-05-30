package twitch

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/merger3/camserver/managers/alias"
	"github.com/merger3/camserver/managers/cache"

	//lint:file-ignore ST1001 I want to use dot imports
	. "github.com/merger3/camserver/modules/core"

	"github.com/gempir/go-twitch-irc/v4"
)

type ValidationResponse struct {
	ClientID   string   `json:"client_id"`
	Login      string   `json:"login"`
	Scopes     []string `json:"scopes"`
	UserID     string   `json:"user_id"`
	Expiration int64    `json:"expires_in"`
}

type Message struct {
	User twitch.User

	Raw     string
	Type    twitch.MessageType
	RawType string
	Message string
	Channel string

	Original twitch.PrivateMessage
}

func NewMessageFromPrivateMessage(pm twitch.PrivateMessage) Message {
	return Message{
		User:     pm.User,
		Raw:      pm.Raw,
		Type:     pm.Type,
		RawType:  pm.RawType,
		Message:  pm.Message,
		Channel:  pm.Channel,
		Original: pm,
	}
}

func NewMessageFromUserStateMessage(pm twitch.UserStateMessage) Message {
	return Message{
		User:    pm.User,
		Raw:     pm.Raw,
		Type:    pm.Type,
		RawType: pm.RawType,
		Message: pm.Message,
		Channel: pm.Channel,
	}
}

type Listener func(message twitch.PrivateMessage)

type User struct {
	Username        string
	Token           string
	Client          *twitch.Client
	ActiveListeners []Listener
	LastMessage     string
	MessageQueue    []Command
	QueueRunning    bool
	APIKey          string
}

func (u *User) CallUsersListeners(message twitch.PrivateMessage) {
	for _, listener := range u.ActiveListeners {
		listener(message)
	}
}

func (u *User) QueueMessage(message Command) {
	u.MessageQueue = append(u.MessageQueue, message)
	ticker := time.NewTicker(1100 * time.Millisecond)

	if !u.QueueRunning {
		u.QueueRunning = true
		go u.RunQueue(ticker)
	}
}

func (u *User) RunQueue(ticker *time.Ticker) {
	processLastMessage := func() {
		lastIndex := len(u.MessageQueue) - 1
		if lastIndex < 0 {
			return
		}

		if !strings.HasPrefix(u.MessageQueue[lastIndex].Command, "!ptzgetcam") {
			lastIndex = 0
		}

		if u.MessageQueue[lastIndex].Command == u.LastMessage {
			u.MessageQueue[lastIndex].Command = fmt.Sprintf("%s .", u.MessageQueue[lastIndex].Command)
		}
		u.LastMessage = u.MessageQueue[lastIndex].Command
		fmt.Printf("[%v] <%v>: %v\n", time.Now().Format("03:04 PM"), u.Username, u.MessageQueue[lastIndex].Command)
		u.Client.Say(u.MessageQueue[lastIndex].Channel, u.MessageQueue[lastIndex].Command)
		u.MessageQueue = slices.Delete(u.MessageQueue, lastIndex, lastIndex+1)
	}

	processLastMessage()

	for range ticker.C {
		if len(u.MessageQueue) == 0 {
			ticker.Stop()
			u.QueueRunning = false
			return
		}
		processLastMessage()
	}
}

type TwitchManager struct {
	Clients    map[string]*User
	Users      map[string]string
	Cache      *cache.CacheManager
	Aliases    *alias.AliasManager
	HTTPClient *http.Client
	OAuth      OAuthTokenManager
	AuthMap    map[string]bool
	Channel    string
	Sentinel   string
	APIKey     string
	Listeners  map[string]Listener
}

func (tm *TwitchManager) SendAPIMessage(message Command) (http.Response, error) {
	fmt.Printf("Sending API command: %+v\n", message)
	url := "https://api.ptz.app:2053/api/command"

	requestBody, err := json.Marshal(Payload{Message: message.Command})
	if err != nil {
		return http.Response{}, err
	}

	request, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return http.Response{}, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tm.APIKey))
	request.Header.Set("Content-Type", "application/json")

	rsp, err := tm.HTTPClient.Do(request)
	if err != nil {
		return http.Response{}, err
	}

	return *rsp, nil
}

func NewTwitchManager(channel, sentinel, token string, cache *cache.CacheManager, aliases *alias.AliasManager, client *http.Client) *TwitchManager {
	tm := TwitchManager{Clients: make(map[string]*User), Users: make(map[string]string), Cache: cache, OAuth: NewOAuthTokenManager(), Channel: channel, Sentinel: sentinel, APIKey: token, Aliases: aliases, HTTPClient: client, Listeners: make(map[string]Listener)}
	tm.CreateListeners()
	tm.AuthMap = createAuthMap()
	tm.AddClient("merger4", tm.OAuth.AccessToken, []Listener{tm.Listeners["scenecams"], tm.Listeners["ptzlist"], tm.Listeners["swap"], tm.Listeners["botSwap"], tm.Listeners["resync"], tm.Listeners["misswap"]})
	tm.Clients["merger4"].Client.Join("alveussanctuary")

	return &tm
}

func (tm *TwitchManager) CreateListeners() {

	botSwapRE := regexp.MustCompile(`^\w+: Swap (\w+ \w+) ?`)

	tm.Listeners["scenecams"] = func(message twitch.PrivateMessage) {
		if match, _ := regexp.MatchString(`^Scene: \w+ Current Scene: \w+ Cams: ((\d: \w+,? ?)+)$`, message.Message); tm.Cache != nil && message.User.Name == tm.Sentinel && match {
			tm.Cache.ParseScene(message.Message)
		}
	}

	tm.Listeners["ptzlist"] = func(message twitch.PrivateMessage) {
		if match, _ := regexp.MatchString(`^Current Scene: \w+ Cams: ((\d: \w+,? ?)+)$`, message.Message); tm.Cache != nil && message.User.Name == tm.Sentinel && match {
			tm.Cache.ParseScene(message.Message)
		}
	}

	tm.Listeners["swap"] = func(message twitch.PrivateMessage) {
		if match, _ := regexp.MatchString(`^\!swap \w+ \w+ ?`, message.Message); tm.Cache != nil && match && tm.CheckUsername(message.User.Name) {
			args := strings.Split(message.Message, " ")[1:]
			tm.Cache.ProcessSwap(args[0], args[1])
			fmt.Printf("%v\n", tm.Cache.Cams)
		}
	}

	tm.Listeners["botSwap"] = func(message twitch.PrivateMessage) {
		if match, _ := regexp.MatchString(`^\w+: Swap (\w+ \w+) ?`, message.Message); tm.Cache != nil && message.User.Name == tm.Sentinel && match {
			matches := botSwapRE.FindStringSubmatch(message.Message)
			args := strings.Split(matches[1], " ")
			tm.Cache.ProcessSwap(args[0], args[1])
			fmt.Printf("%v\n", tm.Cache.Cams)
		}
	}

	tm.Listeners["resync"] = func(message twitch.PrivateMessage) {
		if tm.Cache != nil && (message.Message == "!nightcams" || message.Message == "!livecams") {
			fmt.Println("Invalidating cache")
			tm.Cache.Invalidate()
		}
	}

	tm.Listeners["misswap"] = func(message twitch.PrivateMessage) {
		if tm.Cache != nil && message.User.Name == tm.Sentinel && (message.Message == "Invalid Access" || message.Message == "Invalid Command") {
			fmt.Println("Invalidating cache")
			tm.Cache.Invalidate()
		}
	}

}

func (tm *TwitchManager) AddClient(username, oauth string, listeners []Listener) {
	user := &User{Username: username, Token: oauth, Client: twitch.NewClient(username, fmt.Sprintf("oauth:%s", oauth))}
	user.Client.OnConnect(func() {
		fmt.Printf("Connected %s to Twitch chat\n", username)
	})

	user.Client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		user.CallUsersListeners(message)
	})

	user.ActiveListeners = listeners

	user.Client.Join(tm.Channel)

	tm.Clients[username] = user
	tm.Users[oauth] = username

	go tm.Clients[username].Client.Connect()
}

func (tm TwitchManager) ConnectClients() {
	for _, c := range tm.Clients {
		c.Client.Disconnect()
		go c.Client.Connect()
	}
}

func (tm TwitchManager) ConnectClient(user string) {
	client, ok := tm.Clients[user]
	if !ok {
		return
	}
	client.Client.Disconnect()
	go client.Client.Connect()
}

func (tm TwitchManager) Send(cmd Command) {
	fmt.Println("sending")
	if cmd.User == "merger4" {
		return
	}

	if cmd.Channel == "" {
		cmd.Channel = tm.Channel
	}

	user, ok := tm.Clients[cmd.User]
	if !ok {
		return
	}

	// cmd.Command = strings.ReplaceAll(cmd.Command, "wolfswitch", "wolfindoor")
	if user.Username == "merger3" && !cmd.UseChat && false {
		if cmd.Command == "!scenecams" {
			fmt.Println("Syncing to api")
			err := tm.Cache.SyncCache()
			if errors.Is(err, ErrFailedToSyncCacheWithAPI) {
				fmt.Println("failed to sync to api")
				user.QueueMessage(cmd)
			}
		} else {
			tm.SendAPIMessage(cmd)
		}
	} else {
		user.QueueMessage(cmd)
	}
}

func (tm TwitchManager) GetClickedCam(rect Geom) ClickedCam {
	// return ClickedCam{Found: true, Name: "pasture", Position: 2}
	ch := make(chan string)
	tm.Clients[rect.User].Client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		tm.Clients[rect.User].CallUsersListeners(message)
		if match, _ := regexp.MatchString(`{"cam":"\w+","position":[1-6]}`, message.Message); message.User.Name == tm.Sentinel && match {
			ch <- message.Message
		}
	})

	x, y := rect.GetScaledCoordinates(rect.GetMidpoint())

	tm.Send(Command{User: rect.User, Command: fmt.Sprintf("!ptzgetcam %d %d json", int(math.Round(x)), int(math.Round(y)))})

	var timeout bool
	var cam string
	select {
	case v := <-ch:
		fmt.Println(v)
		cam = v
		timeout = false
		break
	case <-time.After(4 * time.Second):
		timeout = true
		return ClickedCam{}
	}

	tm.Clients[rect.User].Client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		tm.Clients[rect.User].CallUsersListeners(message)
	})

	if timeout {
		return ClickedCam{}
	}

	resp := ClickedCam{Found: true}

	err := json.Unmarshal([]byte(cam), &resp)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return ClickedCam{}
	}

	resp.Name = tm.Aliases.ToCommon(resp.Name)
	return resp
}

func (tm TwitchManager) GetUserFromToken(token string) string {
	user, ok := tm.Users[token]
	if ok {
		fmt.Println("found user in cache")
		return user
	}

	fmt.Println("Going to twitch servers")
	req, _ := http.NewRequest(http.MethodGet, "https://id.twitch.tv/oauth2/validate", nil)
	req.Header.Add("Authorization", fmt.Sprintf("OAuth %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return ""
	}

	if resp.StatusCode == 401 {
		return ""
	}

	validation := ValidationResponse{}
	var b []byte
	if b, err = io.ReadAll(resp.Body); err != nil {
		return ""
	}

	json.Unmarshal(b, &validation)

	return validation.Login
}

// This function used to be used to improve cache consistency by ignoring actions done by people without perms, reduce hard resyncs
// Since all subs have perms now it no longer makes sense to maintain
func (tm TwitchManager) CheckUsername(username string) bool {
	return true
	// _, ok := tm.AuthMap[username]
	// if !ok {
	// 	return false
	// } else {
	// 	return true
	// }
}

func createAuthMap() map[string]bool {

	// This should be offloaded to a config file
	commandAdmins := []string{"spacevoyage", "maya", "theconnorobrien", "alveussanctuary"}
	commandSuperUsers := []string{"ellaandalex", "dionysus1911", "dannydv", "maxzillajr", "illjx", "kayla_alveus",
		"alex_b_patrick", "lindsay_alveus", "strickknine", "tarantulizer", "spiderdaynightlive",
		"srutiloops", "evantomology", "amanda2815"}
	commandMods := []string{"pjeweb", "loganrx_", "mattipv4", "mik_mwp", "96allskills", "wazix11"}
	commandOperator := []string{"96allskills", "stolenarmy_", "berlac", "dansza", "loganrx_", "merger3", "nitelitedf",
		"purplemartinconservation", "lazygoosepxls", "alxiszzz", "shutupleonard",
		"taizun", "lumberaxe1", "glennvde", "wolfone_", "dohregard", "lakel1", "darkrow_",
		"minipurrl", "gnomechildboi", "danman149", "hunnybeehelen", "strangecyan",
		"casualruffian", "viphippo", "bagel_deficient", "jugglingdoh", "catonascreen", "sidmaxwell10"}
	commandVips := []string{"tfries_", "sivvii_", "ghandii_", "axialmars", "jazz_peru", "stealfydoge",
		"xano218", "experimentalcyborg", "klav___", "monkarooo", "nixxform", "madcharliekelly",
		"josh_raiden", "jateu", "storesE6", "rebecca_h9", "matthewde", "user_11_11", "huniebeexd",
		"kurtyykins", "breacherman", "bryceisrightjr", "sumaxu", "mariemellie", "ewok_626",
		"quokka64", "otsargh", "likethecheesebri", "just_some_donkus", "fiveacross", "itszalndrin",
		"nicoleeverleigh", "fishymeep", "ponchobee", "nov1cegg, ohnonicoleio"}

	// Create the map and add all users from the lists
	userMap := make(map[string]bool)

	for _, admin := range commandAdmins {
		userMap[admin] = true
	}

	for _, superUser := range commandSuperUsers {
		userMap[superUser] = true
	}

	for _, mod := range commandMods {
		userMap[mod] = true
	}

	for _, operator := range commandOperator {
		userMap[operator] = true
	}

	for _, vip := range commandVips {
		userMap[vip] = true
	}

	// Output the map
	return userMap
}
