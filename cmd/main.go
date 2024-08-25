package main

import (
	"fmt"

	"github.com/merger3/camserver/pkg/click"
	"github.com/merger3/camserver/pkg/config"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/labstack/echo"
)

var modules map[string]Module

type Module interface {
	Init(...any)
	RegisterRoutes(*echo.Echo)
}

func LoadModules() {
	modules["config"] = config.NewConfigModule()
}

func main() {

	// Set up twitch client
	client := twitch.NewClient("merger3", "oauth:51esxuzacga63qijrpwczxq95m8ejc")
	client.OnConnect(func() {
		fmt.Println("Connected to twitch chat")
	})
	client.Join("alveusgg")

	go client.Connect()

	e := echo.New()
	e.Static("/", "build")

	e.POST("/click", click.ClickTangle)
	e.POST("/draw", click.DrawTangle)

	e.POST("/send", func(c echo.Context) error {
		return click.SendCommand(c, client)
	})

	e.POST("/getConfig", ConfigManager.GetPresets)

	e.POST("/getClickedCam", func(c echo.Context) error {
		return ConfigManager.GetClickedCamConfig(c, client)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
