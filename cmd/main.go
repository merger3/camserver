package main

import (
	"fmt"
	"net/http"

	"github.com/merger3/camserver/pkg/click"
	conf "github.com/merger3/camserver/pkg/config"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/labstack/echo"
)

var ConfigManager conf.ConfigManager

func main() {

	ConfigManager = conf.NewConfigManager()

	client := twitch.NewClient("merger3", "oauth:51esxuzacga63qijrpwczxq95m8ejc")

	client.OnConnect(func() {
		fmt.Println("Connected to twitch chat")
	})

	client.Join("alveusgg")

	go client.Connect()

	e := echo.New()
	e.Static("/", "build")
	e.GET("/hello", func(c echo.Context) error {
		client.Say("merger3", "123qwe4")
		return c.String(http.StatusOK, "Hello, World!")
	})

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
