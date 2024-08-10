package click

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/labstack/echo"
)

type Rectangle struct {
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
	Width       float64 `json:"width"`
	Height      float64 `json:"height"`
	FrameWidth  float64 `json:"frameWidth"`
	FrameHeight float64 `json:"frameHeight"`
}

type Command struct {
	Channel string
	Command string `query:"command"`
}

type Zone struct {
	MinX    float64
	MaxX    float64
	MinY    float64
	MaxY    float64
	ScaleX  float64
	ScaleY  float64
	OffsetX float64
	OffsetY float64
	Zone    int
}

var (
	videoWidth  = float64(1920)
	videoHeight = float64(1080)
)

func ClickTangle(ctx echo.Context) error {
	rect := Rectangle{}

	if err := ctx.Bind(&rect); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	x, y := rect.getScaledCoordinates(rect.getMidpoint())

	intX := int(math.Min(math.Round(x), videoWidth))
	intY := int(math.Min(math.Round(y), videoHeight))

	command := fmt.Sprintf("!ptzclick %d %d 100", intX, intY)

	return ctx.JSON(http.StatusOK, map[string]string{
		"x":       strconv.Itoa(intX),
		"y":       strconv.Itoa(intY),
		"command": command,
	})
}

func DrawTangle(ctx echo.Context) error {
	rect := Rectangle{}

	if err := ctx.Bind(&rect); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	x, y := rect.getScaledCoordinates(rect.getMidpoint())
	w, h := rect.getScaledMeasurments((rect.getMeasurements()))

	intX := int(math.Min(math.Round(x), videoWidth))
	intY := int(math.Min(math.Round(y), videoHeight))

	intW := int(math.Round(w))
	intH := int(math.Round(h))

	command := fmt.Sprintf("!ptzdraw %d %d %d %d", intX, intY, intW, intH)

	fmt.Printf("\n%v\n", command)

	return ctx.JSON(http.StatusOK, map[string]string{
		"x":       strconv.Itoa(intX),
		"y":       strconv.Itoa(intY),
		"command": command,
	})

}

func SendCommand(ctx echo.Context, chatter *twitch.Client) error {
	cmd := Command{Channel: "merger3"}

	if err := ctx.Bind(&cmd); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	// chatter.Say("alveusgg", cmd.Command)
	// chatter.Say(cmd.Channel, cmd.Command)

	return ctx.NoContent(http.StatusOK)
}

// func (r Rectangle) getBottomRight() (float64, float64) {
// 	return r.X + r.Width, r.Y + r.Height
// }

// func (r Rectangle) getTopLeft() (float64, float64) {
// 	return r.X, r.Y
// }

func (r Rectangle) getMidpoint() (float64, float64) {
	xMid := r.X + (r.Width / 2)
	yMid := r.Y + (r.Height / 2)

	return xMid, yMid
}

func (r Rectangle) getMeasurements() (float64, float64) {
	return r.Width, r.Height
}

func (r Rectangle) getScaledCoordinates(xIn, yIn float64) (float64, float64) {
	scaleX := r.FrameWidth / videoWidth
	scaleY := r.FrameHeight / videoHeight

	x := xIn / scaleX
	y := yIn / scaleY

	return x, y
}

func (r Rectangle) getScaledMeasurments(wIn, hIn float64) (float64, float64) {
	scaleX := r.FrameWidth / videoWidth
	scaleY := r.FrameHeight / videoHeight

	w := wIn / scaleX
	h := hIn / scaleY

	return w, h
}
