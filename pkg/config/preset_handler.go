package config

import (
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/merger3/camserver/pkg/core"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/labstack/echo"
)

type PresetRequest struct {
	Cam string `json:"camera"`
}

type PresetResponse struct {
	Found          bool        `json:"found"`
	CamPresetsList *CamPresets `json:"camPresets"`
}

func (c ConfigModule) GetCamPresets(ctx echo.Context, client *twitch.Client) error {
	req := core.Geom{}

	if err := ctx.Bind(&req); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	ch := make(chan string)
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		if message.User.Name == "alveussanctuary" && len(strings.Fields(message.Message)) == 1 {
			ch <- message.Message
		}
	})

	scaleX := req.FrameWidth / core.VideoWidth
	scaleY := req.FrameHeight / core.VideoHeight

	x := req.X / scaleX
	y := req.Y / scaleY

	client.Say("alveusgg", fmt.Sprintf("!ptzgetcam %d %d", int(math.Round(x)), int(math.Round(y))))

	var timeout bool
	var cam string
	select {
	case v := <-ch:
		fmt.Println(v)
		cam = v
		timeout = false
		break
	case <-time.After(10 * time.Second):
		timeout = true
		return ctx.NoContent(http.StatusOK)
	}

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {})

	if timeout {
		return ctx.JSON(http.StatusOK, PresetResponse{Found: false, CamPresetsList: nil})
	}

	for _, presets := range c.Cameras {
		if presets.CamName == cam {
			return ctx.JSON(http.StatusOK, PresetResponse{Found: true, CamPresetsList: &presets})
		}
	}

	return ctx.JSON(http.StatusOK, PresetResponse{Found: false, CamPresetsList: nil})
}
