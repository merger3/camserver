package main

import (
	"crypto/subtle"
	"fmt"
	"net/http"

	"github.com/merger3/camserver/managers/alias"
	"github.com/merger3/camserver/managers/cache"
	"github.com/merger3/camserver/managers/twitch"
	"github.com/merger3/camserver/modules/click"
	"github.com/merger3/camserver/modules/config"
	"github.com/merger3/camserver/modules/core"
	"github.com/merger3/camserver/modules/menu"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	resources map[string]any
	modules   map[string]Module
)

type Module interface {
	Init(map[string]any)
	RegisterRoutes(*echo.Echo)
}

func LoadResources() {
	resources = make(map[string]any)

	resources["aliases"] = *alias.NewAliasManager()
	resources["cache"] = cache.NewCacheManager()

	tm := twitch.NewTwitchManager("merger3", "merger3", resources["cache"].(*cache.CacheManager), resources["aliases"].(alias.AliasManager))
	tm.AddClient("merger3", "oauth:51esxuzacga63qijrpwczxq95m8ejc")
	tm.ConnectClients()
	resources["twitch"] = tm
}

func LoadModules(e *echo.Echo) {
	modules = make(map[string]Module)

	modules["config"] = config.NewConfigModule()
	modules["click"] = click.NewClickModule()
	modules["menu"] = menu.NewMenuModule()

	for _, v := range modules {
		v.Init(resources)
		v.RegisterRoutes(e)
	}
}

func main() {
	LoadResources()

	e := echo.New()
	e.Static("/", "build")

	LoadModules(e)

	e.POST("/send", func(ctx echo.Context) error {
		cmd := core.Command{User: "merger3"}

		if err := ctx.Bind(&cmd); err != nil {
			fmt.Printf("%v\n", err)
			return err
		}

		resources["twitch"].(*twitch.TwitchManager).Send(cmd)

		return ctx.NoContent(http.StatusOK)
	})

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("merger")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("Merger!23")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	// if err := e.StartTLS(":8443", "cert.pem", "cert.key"); err != http.ErrServerClosed {
	// 	e.Logger.Fatal(err)
	// }
	e.Logger.Fatal(e.Start(":1323"))
}
