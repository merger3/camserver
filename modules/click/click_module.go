package click

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/merger3/camserver/managers/alias"
	"github.com/merger3/camserver/managers/cache"
	"github.com/merger3/camserver/managers/twitch"
	"github.com/merger3/camserver/modules/core"
)

type ClickModule struct {
	Twitch  *twitch.TwitchManager
	Cache   *cache.CacheManager
	Aliases *alias.AliasManager
}

func NewClickModule() *ClickModule {
	return &ClickModule{}
}

func (c ClickModule) RegisterRoutes(server *echo.Echo) {
	server.POST("/api/click", c.ClickTangle)
	server.POST("/api/draw", c.DrawTangle)
	server.POST("/api/camera", c.GetCam)
}

func (c *ClickModule) Init(resources map[string]any) {
	c.Twitch = resources["twitch"].(*twitch.TwitchManager)
	c.Cache = resources["cache"].(*cache.CacheManager)
	c.Aliases = resources["aliases"].(*alias.AliasManager)
}

func (c ClickModule) GetCam(ctx echo.Context) error {
	fmt.Println("getting cam")

	req := core.Geom{}

	if err := ctx.Bind(&req); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	if err := (&echo.DefaultBinder{}).BindHeaders(ctx, &req); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	var cam core.ClickedCam
	// Check cache first if the cache is valid
	if c.Cache.IsSynced {
		cam = c.Cache.FetchFromCache(req.Position)
		fmt.Printf("%+v\n", cam)
		if cam.Found {
			cam.HitCache = true
		}
	}

	// Try coordinates and ptzgetcam
	if !cam.Found {
		cam = c.Twitch.GetClickedCam(req)
	}

	// Give up
	if !cam.Found {
		return ctx.JSON(http.StatusNotFound, cam)
	}

	cam.Name = c.Aliases.ToBase(cam.Name)
	return ctx.JSON(http.StatusOK, cam)
}
