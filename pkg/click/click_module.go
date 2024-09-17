package click

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/labstack/echo"
	"github.com/merger3/camserver/pkg/core"
)

var aliases = map[string]string{"parrot": "parrots", "rat": "rat1", "marmoset": "marmout", "crow": "crowin"}

type ClickModule struct {
	Client *twitch.Client
}

func NewClickModule() *ClickModule {
	return &ClickModule{}
}

func (c ClickModule) RegisterRoutes(server *echo.Echo) {
	server.POST("/click", c.ClickTangle)
	server.POST("/draw", c.DrawTangle)
	server.POST("/getcam", c.GetCamFromCoordinates)
}

func (c *ClickModule) Init(resources map[string]any) {
	c.Client = resources["twitch"].(*twitch.Client)
}

type ClickedCam struct {
	Found    bool   `json:"found"`
	Name     string `json:"cam"`
	Position int    `json:"position"`
}

func GetClickedCam(client *twitch.Client, rect core.Geom) ClickedCam {
	// return ClickedCam{Found: true, Name: "pasture", Position: 2}
	ch := make(chan string)
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		if match, _ := regexp.MatchString(`{"cam":"\w+","position":[1-6]}`, message.Message); message.User.Name == "alveussanctuary" && match {
			ch <- message.Message
		}
	})

	x, y := rect.GetScaledCoordinates(rect.GetMidpoint())

	client.Say("alveusgg", fmt.Sprintf("!ptzgetcam %d %d json", int(math.Round(x)), int(math.Round(y))))

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

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {})

	if timeout {
		return ClickedCam{}
	}

	resp := ClickedCam{Found: true}

	err := json.Unmarshal([]byte(cam), &resp)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return ClickedCam{}
	}

	camAlias, ok := aliases[resp.Name]
	if !ok {
		return resp
	} else {
		resp.Name = camAlias
		return resp
	}

}

func (c ClickModule) GetCamFromCoordinates(ctx echo.Context) error {
	req := core.Geom{}

	if err := ctx.Bind(&req); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	cam := GetClickedCam(c.Client, req)

	return ctx.JSON(http.StatusOK, cam)
}
