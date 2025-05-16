package cams

import (
	"github.com/labstack/echo/v4"
	"github.com/merger3/camserver/managers/alias"
	"github.com/merger3/camserver/managers/cache"
	"github.com/merger3/camserver/managers/twitch"
)

type CamsModule struct {
	Twitch  *twitch.TwitchManager
	Cache   *cache.CacheManager
	Aliases *alias.AliasManager
}

func NewCamsModule() *CamsModule {
	return &CamsModule{}
}

func (c CamsModule) RegisterRoutes(server *echo.Echo) {
	server.GET("/api/layout", c.GetLayout)
}

func (c *CamsModule) Init(resources map[string]any) {
	c.Twitch = resources["twitch"].(*twitch.TwitchManager)
	c.Cache = resources["cache"].(*cache.CacheManager)
	c.Aliases = resources["aliases"].(*alias.AliasManager)
}
