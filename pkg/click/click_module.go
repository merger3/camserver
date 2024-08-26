package click

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/labstack/echo"
	"github.com/merger3/camserver/pkg/core"
)

type ClickModule struct {
	Client *twitch.Client
}

func NewConfigModule() *ClickModule {
	return &ClickModule{}
}

func (c ClickModule) RegisterRoutes(server *echo.Echo) {
	server.POST("/click", c.ClickTangle)
	server.POST("/draw", c.DrawTangle)
}

func (c *ClickModule) Init(resources map[string]any) {
	c.Client = resources["twitch"].(*twitch.Client)
}

func GetClickedCam(client *twitch.Client, rect core.Geom) string {
	ch := make(chan string)
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		if message.User.Name == "alveussanctuary" && len(strings.Fields(message.Message)) == 1 {
			ch <- message.Message
		}
	})

	x, y := rect.GetScaledCoordinates(rect.GetMidpoint())

	client.Say("alveusgg", fmt.Sprintf("!ptzgetcam %d %d", int(math.Round(x)), int(math.Round(y))))

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
		return ""
	}

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {})

	if timeout {
		return ""
	}

	return cam
}
