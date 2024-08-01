package main

import (
	"fmt"
	"net/http"

	"github.com/merger3/camserver/pkg/click"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/labstack/echo"
)

func main() {

	client := twitch.NewClient("merger3", "oauth:51esxuzacga63qijrpwczxq95m8ejc")

	client.OnConnect(func() {
		fmt.Println("Connected to twitch chat")
	})

	client.Join("merger3")

	go client.Connect()

	e := echo.New()
	e.Static("/", "build")
	e.GET("/hello", func(c echo.Context) error {
		client.Say("merger3", "123qwe4")
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/draw", func(c echo.Context) error {
		return click.DrawTangle(c, client)
	})

	// e.POST("/draw", click.DrawTangle)
	// e.GET("/hello", func(c echo.Context) error {
	// 	client.Say("merger3", "123qwe4")
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	e.Logger.Fatal(e.Start(":1323"))

	// chat.Client.Say("merger3", "hellos world!")
}
