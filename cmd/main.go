package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/merger3/camserver/pkg/click"
	"github.com/merger3/camserver/pkg/config"
	"github.com/merger3/camserver/pkg/core"
	"github.com/merger3/camserver/pkg/menu"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/labstack/echo"
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
			cmd.Command = fmt.Sprintf("%s 4", cmd.Command)
		}
		resources["twitch"].(*twitch.Client).Say(cmd.Channel, cmd.Command)

		return ctx.NoContent(http.StatusOK)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
