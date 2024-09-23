package click

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/merger3/camserver/modules/core"

	// "github.com/gempir/go-twitch-irc/v4"
	"github.com/labstack/echo/v4"
)

func (c ClickModule) ClickTangle(ctx echo.Context) error {
	rect := core.Geom{}

	if err := ctx.Bind(&rect); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	x, y := rect.GetScaledCoordinates(rect.GetMidpoint())

	intX := int(math.Min(math.Round(x), core.VideoWidth))
	intY := int(math.Min(math.Round(y), core.VideoHeight))

	command := fmt.Sprintf("!ptzclick %d %d 100", intX, intY)

	return ctx.JSON(http.StatusOK, map[string]string{
		"x":       strconv.Itoa(intX),
		"y":       strconv.Itoa(intY),
		"command": command,
	})
}

func (c ClickModule) DrawTangle(ctx echo.Context) error {
	rect := core.Geom{}

	if err := ctx.Bind(&rect); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	x, y := rect.GetScaledCoordinates(rect.GetTopLeft())
	w, h := rect.GetScaledMeasurments((rect.GetMeasurements()))

	intX := int(math.Min(math.Round(x), core.VideoWidth))
	intY := int(math.Min(math.Round(y), core.VideoHeight))

	intW := int(math.Round(w))
	intH := int(math.Round(h))

	command := fmt.Sprintf("!ptzdraw %d %d %d %d", intX, intY, intW, intH)

	return ctx.JSON(http.StatusOK, map[string]string{
		"x":       strconv.Itoa(intX),
		"y":       strconv.Itoa(intY),
		"command": command,
	})

}
