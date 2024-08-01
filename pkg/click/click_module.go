package click

import (
	"fmt"
	"math"
	"net/http"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/labstack/echo"
)

type Rectangle struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func DrawTangle(ctx echo.Context, chatter *twitch.Client) error {
	rect := Rectangle{}

	if err := ctx.Bind(&rect); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	x := math.Round(rect.X + (rect.Width / 2))
	y := math.Round(rect.Y + (rect.Height / 2))

	chatter.Say("merger3", "ad;ljiowej")

	return ctx.JSON(http.StatusOK, map[string]float64{
		"x": x,
		"y": y,
	})
}
