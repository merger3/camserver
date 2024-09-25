package twitch

import (
	"encoding/json"
	"fmt"
	"math"
	"regexp"
	"strings"
	"time"

	"github.com/merger3/camserver/managers/alias"
	"github.com/merger3/camserver/managers/cache"
	. "github.com/merger3/camserver/modules/core"

	"github.com/gempir/go-twitch-irc/v4"
)

type Listener func(message twitch.PrivateMessage, user string)

type User struct {
	Username        string
	Client          *twitch.Client
	ActiveListeners []Listener
}

type TwitchManager struct {
	Clients   map[string]*User
	Cache     *cache.CacheManager
	Aliases   alias.AliasManager
	Channel   string
	Sentinel  string
	Listeners map[string]Listener
}

func NewTwitchManager(channel, sentinel string, cache *cache.CacheManager, aliases alias.AliasManager) *TwitchManager {
	tm := TwitchManager{Clients: make(map[string]*User), Cache: cache, Channel: channel, Sentinel: sentinel, Aliases: aliases, Listeners: make(map[string]Listener)}
	tm.CreateListeners()
	return &tm
}

func (tm *TwitchManager) CreateListeners() {
	tm.Listeners["scenecams"] = func(message twitch.PrivateMessage, user string) {
		if match, _ := regexp.MatchString(`^1: \w+, 2: \w+, 3: \w+, 4: \w+, 5: \w+, 6: \w+$`, message.Message); tm.Cache != nil && message.User.Name == tm.Sentinel && match {
			fmt.Println("Scenecams Match")
			tm.Cache.ParseScene(message.Message)
		}
	}

	tm.Listeners["swap"] = func(message twitch.PrivateMessage, user string) {
		if match, _ := regexp.MatchString(`^\!swap \w+ \w+$`, message.Message); tm.Cache != nil && match {
			fmt.Println("Swap Match")
			fmt.Printf("%v\n", tm.Cache.Cams)
			args := strings.Split(message.Message, " ")[1:]
			tm.Cache.ProcessSwap(args[0], args[1])
			fmt.Printf("%v\n", tm.Cache.Cams)
		}
	}

	tm.Listeners["resync"] = func(message twitch.PrivateMessage, user string) {
		if tm.Cache != nil && (message.Message == "!nightcams" || message.Message == "!livecams") {
			fmt.Println("Resyncing")
			time.Sleep(1000)
			tm.Send(Command{User: user, Command: "!scenecams"})
		}
	}

}

func (tm *TwitchManager) AddClient(username, oauth string) {
	user := &User{Username: username, Client: twitch.NewClient(username, oauth)}
	user.Client.OnConnect(func() {
		fmt.Printf("Connected %s to Twitch chat\n", username)
	})

	user.Client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		for _, listener := range user.ActiveListeners {
			fmt.Println("Calling Listner")
			listener(message, user.Username)
		}
	})

	user.ActiveListeners = []Listener{tm.Listeners["scenecams"], tm.Listeners["swap"], tm.Listeners["resync"]}

	user.Client.Join(tm.Channel)

	tm.Clients[username] = user
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

	tm.Clients[cmd.User].Client.Say(cmd.Channel, cmd.Command)
}

func (tm TwitchManager) GetClickedCam(rect Geom) ClickedCam {
	// return ClickedCam{Found: true, Name: "pasture", Position: 2}
	ch := make(chan string)
	tm.Clients[rect.User].Client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		if match, _ := regexp.MatchString(`{"cam":"\w+","position":[1-6]}`, message.Message); message.User.Name == "alveussanctuary" && match {
			ch <- message.Message
		}
	})

	x, y := rect.GetScaledCoordinates(rect.GetMidpoint())

	tm.Clients[rect.User].Client.Say(tm.Channel, fmt.Sprintf("!ptzgetcam %d %d json", int(math.Round(x)), int(math.Round(y))))

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

	tm.Clients[rect.User].Client.OnPrivateMessage(func(message twitch.PrivateMessage) {})

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
