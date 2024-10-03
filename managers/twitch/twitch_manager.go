package twitch

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"regexp"
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
}

func NewMessageFromPrivateMessage(pm twitch.PrivateMessage) Message {
	return Message{
		User:    pm.User,
		Raw:     pm.Raw,
		Type:    pm.Type,
		RawType: pm.RawType,
		Message: pm.Message,
		Channel: pm.Channel,
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

type Listener func(message Message, user string)

type User struct {
	Username        string
	Client          *twitch.Client
	ActiveListeners []Listener
	LastMessage     string
}

type TwitchManager struct {
	Clients   map[string]*User
	Cache     *cache.CacheManager
	Aliases   alias.AliasManager
	OAuth     OAuthTokenManager
	Channel   string
	Sentinel  string
	Listeners map[string]Listener
}

func NewTwitchManager(channel, sentinel string, cache *cache.CacheManager, aliases alias.AliasManager) *TwitchManager {
	tm := TwitchManager{Clients: make(map[string]*User), Cache: cache, OAuth: NewOAuthTokenManager(), Channel: channel, Sentinel: sentinel, Aliases: aliases, Listeners: make(map[string]Listener)}
	tm.CreateListeners()

	tm.AddClient("merger4", tm.OAuth.AccessToken, []Listener{tm.Listeners["scenecams"], tm.Listeners["swap"], tm.Listeners["resync"], tm.Listeners["misswap"]})
	return &tm
}

func (tm *TwitchManager) CreateListeners() {
	tm.Listeners["scenecams"] = func(message Message, user string) {
		if match, _ := regexp.MatchString(`^1: \w+, 2: \w+, 3: \w+, 4: \w+, 5: \w+, 6: \w+$`, message.Message); tm.Cache != nil && message.User.Name == tm.Sentinel && match {
			tm.Cache.ParseScene(message.Message)
		}
	}

	tm.Listeners["swap"] = func(message Message, user string) {
		if match, _ := regexp.MatchString(`^\!swap \w+ \w+$`, message.Message); tm.Cache != nil && match {
			args := strings.Split(message.Message, " ")[1:]
			tm.Cache.ProcessSwap(args[0], args[1])
			fmt.Printf("%v\n", tm.Cache.Cams)
		}
	}

	tm.Listeners["resync"] = func(message Message, user string) {
		if tm.Cache != nil && (message.Message == "!nightcams" || message.Message == "!livecams") {
			fmt.Println("Invalidating cache")
			tm.Cache.Invalidate()
		}
	}

	tm.Listeners["misswap"] = func(message Message, user string) {
		if tm.Cache != nil && message.User.Name == tm.Sentinel && (message.Message == "Invalid Access" || message.Message == "Invalid Command") {
			fmt.Println("Invalidating cache")
			tm.Cache.Invalidate()
		}
	}

}

func (u *User) CallUsersListeners(message Message) {
	for _, listener := range u.ActiveListeners {
		listener(message, u.Username)
	}
}

func (tm *TwitchManager) AddClient(username, oauth string, listeners []Listener) {
	user := &User{Username: username, Client: twitch.NewClient(username, fmt.Sprintf("oauth:%s", oauth))}
	user.Client.OnConnect(func() {
		fmt.Printf("Connected %s to Twitch chat\n", username)
	})

	user.Client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		user.CallUsersListeners(NewMessageFromPrivateMessage(message))
	})

	user.ActiveListeners = listeners

	user.Client.Join(tm.Channel)

	tm.Clients[username] = user

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
	if cmd.Channel == "" {
		cmd.Channel = tm.Channel
	}

	user := tm.Clients[cmd.User]
	if cmd.Command == user.LastMessage {
		cmd.Command = fmt.Sprintf("%s .", cmd.Command)
	}
	user.Client.Say(cmd.Channel, cmd.Command)
	user.LastMessage = cmd.Command
}

func (tm TwitchManager) GetClickedCam(rect Geom) ClickedCam {
	// return ClickedCam{Found: true, Name: "pasture", Position: 2}

	ch := make(chan string)
	tm.Clients[rect.User].Client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		tm.Clients[rect.User].CallUsersListeners(NewMessageFromPrivateMessage(message))

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
	case <-time.After(10 * time.Second):
		timeout = true
		return ClickedCam{}
	}

	tm.Clients[rect.User].Client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		tm.Clients[rect.User].CallUsersListeners(NewMessageFromPrivateMessage(message))
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

	// resp.Name = tm.Aliases.ToCommon(resp.Name)
	return resp
}

func (tm TwitchManager) GetUserFromToken(token string) string {
	req, _ := http.NewRequest(http.MethodGet, "https://id.twitch.tv/oauth2/validate", nil)
	req.Header.Add("Authorization", fmt.Sprintf("OAuth %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}

	validation := ValidationResponse{}
	var b []byte
	if b, err = io.ReadAll(resp.Body); err != nil {
		return err.Error()
	}

	json.Unmarshal(b, &validation)

	return validation.Login
}
