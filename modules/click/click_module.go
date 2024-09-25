package click

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/merger3/camserver/managers/twitch"
	"github.com/merger3/camserver/modules/core"
)

var aliases = map[string]string{"parrot": "parrots", "rat": "rat1", "marmoset": "marmout", "crow": "crowin"}

type ClickModule struct {
	Twitch *twitch.TwitchManager
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
	c.Twitch = resources["twitch"].(*twitch.TwitchManager)
}

func (c ClickModule) GetCamFromCoordinates(ctx echo.Context) error {
	req := core.Geom{}

	if err := ctx.Bind(&req); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	cam := c.Twitch.GetClickedCam(req)

	return ctx.JSON(http.StatusOK, cam)
}
