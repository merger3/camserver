package twitch

//lint:file-ignore ST1001 I want to use dot imports
import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v4"
)

type TwitchClient struct {
	client  *twitch.Client
	channel string
}

func NewTwitchClient(botUsername, oauthToken, channel string) *TwitchClient {
	client := twitch.NewClient(botUsername, oauthToken)
	return &TwitchClient{
		client:  client,
		channel: channel,
	}
}

func (t *TwitchClient) Connect() error {
	t.client.Join(t.channel)

	t.client.OnConnect(func() {
		fmt.Println("Connected to Twitch chat")
	})

	t.client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Printf("Received message from %s: %s\n", message.User.DisplayName, message.Message)
	})

	go func() {
		err := t.client.Connect()
		if err != nil {
			fmt.Printf("Error connecting to Twitch: %v", err)
		}
	}()

	return nil
}

func (t *TwitchClient) SendMessage(message string) {
	t.client.Say(t.channel, message)
}
