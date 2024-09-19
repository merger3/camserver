package main

import (
	"crypto/subtle"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/merger3/camserver/pkg/click"
	"github.com/merger3/camserver/pkg/config"
	"github.com/merger3/camserver/pkg/core"
	"github.com/merger3/camserver/pkg/menu"

	"github.com/gempir/go-twitch-irc/v4"
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

	// Set up twitch client
	client := twitch.NewClient("merger3", "oauth:51esxuzacga63qijrpwczxq95m8ejc")
	client.OnConnect(func() {
		fmt.Println("Connected to twitch chat")
	})
	client.Join("alveusgg")
	resources["twitch"] = client
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

	go resources["twitch"].(*twitch.Client).Connect()

	e := echo.New()
	e.Static("/", "build")

	LoadModules(e)

	e.POST("/send", func(ctx echo.Context) error {
		cmd := core.Command{Channel: "alveusgg"}

		if err := ctx.Bind(&cmd); err != nil {
			fmt.Printf("%v\n", err)
			return err
		}

		if strings.HasPrefix(cmd.Command, "!ptzdraw") {
			cmd.Command = fmt.Sprintf("%s 5", cmd.Command)
		}
		resources["twitch"].(*twitch.Client).Say(cmd.Channel, cmd.Command)

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

	e.GET("/proxy", func(c echo.Context) error {
		url := c.QueryParam("url")

		// Check if the URL starts with "http://"
		if url == "" || !startsWithHTTP(url) {
			return c.String(http.StatusBadRequest, "Invalid URL. Make sure it starts with http://")
		}

		// Fetch the content from the external HTTP source
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Error fetching URL:", err)
			return c.String(http.StatusInternalServerError, "Failed to fetch the content")
		}
		defer resp.Body.Close()

		// Stream the fetched content to the client
		c.Response().Header().Set("Content-Type", resp.Header.Get("Content-Type"))
		c.Response().WriteHeader(resp.StatusCode)
		_, err = io.Copy(c.Response().Writer, resp.Body)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error streaming content")
		}
		return nil
	})

	if err := e.StartTLS(":8443", "cert.pem", "cert.key"); err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}
	// e.Logger.Fatal(e.Start(":8443"))
}

// Utility function to check if the URL starts with http://
func startsWithHTTP(url string) bool {
	return len(url) > 7 && url[:7] == "http://"
}
