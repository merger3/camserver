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

	w1 = videoWidth / 3
	w2 = videoWidth / 1.5

	h1 = videoHeight / 3
	h2 = videoHeight / 1.5

	zones = []Zone{
		{MinX: w1, MaxX: videoWidth, MinY: 0, MaxY: h2, ScaleX: 1.5, ScaleY: 1.5, OffsetX: w1, OffsetY: 0, Zone: 1},
		{MinX: 0, MaxX: w1, MinY: 0, MaxY: h1, ScaleX: 3, ScaleY: 3, OffsetX: 0, OffsetY: 0, Zone: 2},
		{MinX: 0, MaxX: w1, MinY: h1, MaxY: h2, ScaleX: 3, ScaleY: 3, OffsetX: 0, OffsetY: h1, Zone: 3},
		{MinX: 0, MaxX: w1, MinY: h2, MaxY: videoHeight, ScaleX: 3, ScaleY: 3, OffsetX: 0, OffsetY: h2, Zone: 4},
		{MinX: w1, MaxX: w2, MinY: h2, MaxY: videoHeight, ScaleX: 3, ScaleY: 3, OffsetX: w1, OffsetY: h2, Zone: 5},
		{MinX: w2, MaxX: videoWidth, MinY: h2, MaxY: videoHeight, ScaleX: 3, ScaleY: 3, OffsetX: w2, OffsetY: h2, Zone: 6},
	}
)

func ClickTangle(ctx echo.Context) error {
	rect := Rectangle{}

	if err := ctx.Bind(&rect); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	x, y := rect.getScaledCoordinates(rect.getMidpoint())

	intX := int(math.Round(x))
	intY := int(math.Round(y))

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

	x1, y1 := getRelativeCoordinates(rect.getScaledCoordinates(rect.getTopLeft()))
	x2, y2 := getRelativeCoordinates(rect.getScaledCoordinates(rect.getBottomRight()))

	zoom := findZoom(x1, y1, x2, y2)

	intX := int(math.Round(x))
	intY := int(math.Round(y))

	command := fmt.Sprintf("!ptzclick %d %d %d", intX, intY, zoom)

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

	chatter.Say(cmd.Channel, cmd.Command)

	return ctx.NoContent(http.StatusOK)
}

func (r Rectangle) getBottomRight() (float64, float64) {
	return r.X + r.Width, r.Y + r.Height
}

func (r Rectangle) getTopLeft() (float64, float64) {
	return r.X, r.Y
}

func (r Rectangle) getMidpoint() (float64, float64) {
	scaledX := r.X + (r.Width / 2)
	scaledY := r.Y + (r.Height / 2)

	return scaledX, scaledY
}

func (r Rectangle) getScaledCoordinates(xIn, yIn float64) (float64, float64) {
	scaleX := r.FrameWidth / videoWidth
	scaleY := r.FrameHeight / videoHeight

	x := xIn / scaleX
	y := yIn / scaleY

	return x, y
}

func getRelativeCoordinates(x, y float64) (float64, float64) {
	var xUnscaled, yUnscaled float64
	for _, z := range zones {
		if (xUnscaled > z.MinX && xUnscaled < z.MaxX) && (yUnscaled > z.MinY && yUnscaled < z.MaxY) {
			x, y = convertCoordinates(xUnscaled, yUnscaled, z)
			break
		}
	}
	return x, y
}

func convertCoordinates(x_raw, y_raw float64, zone Zone) (float64, float64) {
	x := (x_raw - zone.OffsetX) * zone.ScaleX
	y := (y_raw - zone.OffsetY) * zone.ScaleY
	return x, y
}

func findZoom(x1, y1, x2, y2 float64) int {
	box_w := x2 - x1
	box_h := y2 - y1

	zoomWidth := videoWidth / box_w
	zoomHeight := videoHeight / box_h
	zoom := math.Min(zoomWidth, zoomHeight)
	if zoom < 4 {
		zoom = zoom * 100
	} else {
		zoom = zoom * 70

	}

	return int(math.Round(zoom))
}
