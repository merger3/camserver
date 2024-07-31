package twitch

//lint:file-ignore ST1001 I want to use dot imports
import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v4"
)

func main() {
	// or client := twitch.NewAnonymousClient() for an anonymous user (no write capabilities)
	client := twitch.NewClient("yourtwitchusername", "oauth:123123123")

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Message)
	})

	client.Join("merger3")

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
