package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/merger3/camserver/modules/cams"
	"github.com/merger3/camserver/modules/click"
	"github.com/merger3/camserver/modules/config"
	"github.com/merger3/camserver/modules/core"
	"github.com/merger3/camserver/modules/menu"

	"github.com/merger3/camserver/managers/alias"
	"github.com/merger3/camserver/managers/cache"
	"github.com/merger3/camserver/managers/notifications"
	"github.com/merger3/camserver/managers/twitch"

	"github.com/labstack/echo/v4"
)

var (
	resources map[string]any
	modules   map[string]Module
)

type SetupVars struct {
	Channel        string `json:"channel"`
	Sentinel       string `json:"sentinel"`
	APIKey         string `json:"token"`
	PushoverKey    string `json:"pushover_api_key"`
	PushoverDevice string `json:"pushover_device_key"`
}

type Module interface {
	Init(map[string]any)
	RegisterRoutes(*echo.Echo)
}

func LoadResources() {
	file, err := os.Open(filepath.Join("configs", "setup.json"))
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	var setup SetupVars

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&setup); err != nil {
		log.Fatalf("error unmarshalling JSON: %s", err)
	}

	resources = make(map[string]any)

	resources["token"] = setup.APIKey
	resources["client"] = NewHTTPClient()
	resources["aliases"] = *alias.NewAliasManager()
	resources["cache"] = cache.NewCacheManager(resources["client"].(*http.Client), resources["token"].(string))
	resources["twitch"] = twitch.NewTwitchManager(setup.Channel, setup.Sentinel, resources["token"].(string), resources["cache"].(*cache.CacheManager), resources["aliases"].(alias.AliasManager), resources["client"].(*http.Client))
	resources["notifications"] = notifications.NewNotificationsManager(resources["twitch"].(*twitch.TwitchManager), resources["twitch"].(*twitch.TwitchManager).Clients["merger4"], setup.PushoverKey, setup.PushoverDevice)
}

func LoadModules(e *echo.Echo) {
	modules = make(map[string]Module)

	modules["config"] = config.NewConfigModule()
	modules["click"] = click.NewClickModule()
	modules["cams"] = cams.NewCamsModule()
	modules["menu"] = menu.NewMenuModule()

	for _, v := range modules {
		v.Init(resources)
		v.RegisterRoutes(e)
	}
}

func main() {
	LoadResources()

	e := echo.New()
	LoadModules(e)

	e.POST("/api/send", func(ctx echo.Context) error {
		cmd := core.Command{}

		if err := ctx.Bind(&cmd); err != nil {
			fmt.Printf("%v\n", err)
			return err
		}

		if err := (&echo.DefaultBinder{}).BindHeaders(ctx, &cmd); err != nil {
			fmt.Printf("%v\n", err)
			return err
		}

		resources["twitch"].(*twitch.TwitchManager).Send(cmd)

		return ctx.NoContent(http.StatusOK)
	})

	e.Use(ProcessUser(resources["twitch"].(*twitch.TwitchManager)))
	e.Use(CheckCache(resources["cache"].(*cache.CacheManager), resources["twitch"].(*twitch.TwitchManager)))

	e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}

func NewHTTPClient() *http.Client {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // Timeout for establishing connection
			KeepAlive: 30 * time.Second, // Keep-alive period
		}).DialContext,
		MaxIdleConns:          100,              // Maximum number of idle connections
		IdleConnTimeout:       90 * time.Second, // Idle timeout
		TLSHandshakeTimeout:   10 * time.Second, // Timeout for TLS handshake
		ExpectContinueTimeout: 1 * time.Second,  // Timeout for Expect: 100-continue responses
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   60 * time.Second, // Overall request timeout
	}

	return client
}

func setCookie(c echo.Context, token string) {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().AddDate(3, 0, 0)
	cookie.MaxAge = 60 * 60 * 24 * 365 * 3
	c.SetCookie(cookie)
}

func checkHeader(tm *twitch.TwitchManager, c echo.Context) (string, error) {
	var token string
	var username string

	token = c.Request().Header.Get("X-Twitch-Token")
	if token == "" {
		return "", c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Required X-Twitch-Token header missing",
		})
	} else {
		username = tm.GetUserFromToken(token)
		if username == "" {
			return "", c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Invalid Twitch token sent",
			})
		} else {
			if err := checkPerms(tm, c, username, token, true); err != nil {
				return "", err
			}
		}
	}
	return username, nil
}

func checkPerms(tm *twitch.TwitchManager, c echo.Context, username, token string, shouldSetCookie bool) error {
	if !tm.CheckUsername(username) {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Not authorized",
		})
	} else {
		if shouldSetCookie {
			setCookie(c, token)
		}

		if user, ok := tm.Clients[username]; !ok {
			tm.AddClient(username, token, []twitch.Listener{})
		} else {
			if user.Token != token && tm.GetUserFromToken(user.Token) == "" {
				user.Client.Disconnect()
				tm.AddClient(username, token, []twitch.Listener{})
			}
		}
	}
	return nil
}

func ProcessUser(tm *twitch.TwitchManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var token string
			var username string

			tokenCookie, err := c.Cookie("token")
			if err != nil {
				if username, err = checkHeader(tm, c); err != nil {
					return err
				}
			} else {
				token = tokenCookie.Value
				username = tm.GetUserFromToken(token)
				if username == "" {
					if username, err = checkHeader(tm, c); err != nil {
						return err
					}
				} else {
					if err = checkPerms(tm, c, username, token, false); err != nil {
						return err
					}
				}
			}

			c.Request().Header.Add(core.UsernameHeader, username)
			// fmt.Printf("Twitch Header: %s\n", c.Request().Header.Get(core.UsernameHeader))

			return next(c)
		}
	}
}

func CheckCache(cache *cache.CacheManager, client *twitch.TwitchManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !cache.IsSynced {
				fmt.Printf("Attempting sync at %v, backoff is %v, attempt number %v\n", int(time.Since(cache.LastAttemptedSync).Seconds()), math.Max(6, 6*math.Pow(3, cache.SyncAttempts)), cache.SyncAttempts)
				if time.Since(cache.LastAttemptedSync).Seconds() >= math.Max(6, 6*math.Pow(3, cache.SyncAttempts)) && cache.SyncAttempts <= 6 {
					// client.Send(core.Command{User: c.Request().Header.Get(core.UsernameHeader), Command: "!scenecams"})
					cache.SyncAttempts += 1
					cache.LastAttemptedSync = time.Now()
					err := client.Cache.SyncCache()
					if errors.Is(err, core.ErrFailedToSyncCacheWithAPI) {
						client.Send(core.Command{User: c.Request().Header.Get(core.UsernameHeader), Command: "!scenecams", UseChat: true})
					}
					// client.Send(core.Command{User: c.Request().Header.Get(core.UsernameHeader), Command: "1: toast, 2: parrot, 3: fox, 4: marmoset, 5: wolfden, 6: pasture"})
				}
			}
			return next(c)
		}
	}
}
