package main

import (
	"crypto/subtle"
	"fmt"
	"net/http"
	"time"

	"github.com/merger3/camserver/modules/click"
	"github.com/merger3/camserver/modules/config"
	"github.com/merger3/camserver/modules/core"
	"github.com/merger3/camserver/modules/menu"

	"github.com/merger3/camserver/managers/alias"
	"github.com/merger3/camserver/managers/cache"
	"github.com/merger3/camserver/managers/twitch"

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
	resources["twitch"] = tm
	// tm.AddClient("merger3", "oauth:51esxuzacga63qijrpwczxq95m8ejc")
	// tm.ConnectClients()
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
	e.File("/login", "build/login.html")
	LoadModules(e)

	//str := `https://id.twitch.tv/oauth2/authorize?response_type=token&client_id=s4ouxddi9skb11jriwyzl0ronh1m92&redirect_uri=http://localhost:1323/&scope=user%3Awrite%3Achat`

	e.POST("/send", func(ctx echo.Context) error {
		cmd := core.Command{}

		if err := ctx.Bind(&cmd); err != nil {
			fmt.Printf("%v\n", err)
			return err
		}

		if err := (&echo.DefaultBinder{}).BindHeaders(ctx, &cmd); err != nil {
			fmt.Printf("%v\n", err)
			return err
		}

		fmt.Printf("cmd: %+v\n", cmd)

		resources["twitch"].(*twitch.TwitchManager).Send(cmd)

		return ctx.NoContent(http.StatusOK)
	})

	e.Use(middleware.BasicAuth(func(username string, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("merger3")) == 1 && subtle.ConstantTimeCompare([]byte(password), []byte("Merger!23")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	e.Use(ProcessUser(resources["twitch"].(*twitch.TwitchManager)))
	e.Use(CheckCache(resources["cache"].(*cache.CacheManager), resources["twitch"].(*twitch.TwitchManager)))

	// if err := e.StartTLS(":8443", "cert.pem", "cert.key"); err != http.ErrServerClosed {
	// 	e.Logger.Fatal(err)
	// }
	e.Logger.Fatal(e.Start(":1323"))
}

func ProcessUser(tm *twitch.TwitchManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("X-Twitch-Token")

			user := tm.GetUserFromToken(token)
			fmt.Printf("User: %s\n", user)
			c.Request().Header.Add(core.UsernameHeader, user)
			fmt.Printf("Twitch Header: %s\n", c.Request().Header.Get(core.UsernameHeader))
			if _, ok := tm.Clients[user]; !ok {
				tm.AddClient(user, token)
			}

			return next(c)
		}
	}
}

func CheckCache(cache *cache.CacheManager, client *twitch.TwitchManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Header.Get("X-Twitch-Token") != "" {
				if time.Since(cache.LastSynced).Hours() >= 1 {
					fmt.Printf("Time since: %v", time.Since(cache.LastSynced).Hours())
					fmt.Println("Invalidating cache from middleware")
					cache.Invalidate()
				}
				if !cache.IsSynced {
					client.Send(core.Command{User: c.Request().Header.Get(core.UsernameHeader), Command: "!scenecams"})
				}
			}

			return next(c)
		}
	}
}
