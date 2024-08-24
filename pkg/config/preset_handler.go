package config

import (
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/labstack/echo"
)

var (
	videoWidth  = float64(1920)
	videoHeight = float64(1080)
)

type PresetRequest struct {
	Cam string `json:"camera"`
}

type PresetResponse struct {
	Found          bool        `json:"found"`
	CamPresetsList *CamPresets `json:"camPresets"`
}

func (c ConfigManager) GetCamPresets(ctx echo.Context, client *twitch.Client) error {
	req := ClickedCamRequest{}

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

	scaleX := req.FrameWidth / videoWidth
	scaleY := req.FrameHeight / videoHeight

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

// func (c ConfigManager) GetCamSwaps(ctx echo.Context, client *twitch.Client) error {
// 	req := ClickedCamRequest{}

// 	if err := ctx.Bind(&req); err != nil {
// 		fmt.Printf("%v\n", err)
// 		return err
// 	}

// 	ch := make(chan string)
// 	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
// 		if message.User.Name == "alveussanctuary" && len(strings.Fields(message.Message)) == 1 {
// 			ch <- message.Message
// 		}
// 	})

// 	scaleX := req.FrameWidth / videoWidth
// 	scaleY := req.FrameHeight / videoHeight

// 	x := req.X / scaleX
// 	y := req.Y / scaleY

// 	client.Say("alveusgg", fmt.Sprintf("!ptzgetcam %d %d", int(math.Round(x)), int(math.Round(y))))

// 	var timeout bool
// 	var cam string
// 	select {
// 	case v := <-ch:
// 		fmt.Println(v)
// 		cam = v
// 		timeout = false
// 		break
// 	case <-time.After(10 * time.Second):
// 		timeout = true
// 		return ctx.NoContent(http.StatusOK)
// 	}

// 	client.OnPrivateMessage(func(message twitch.PrivateMessage) {})

// 	if timeout {
// 		return ctx.JSON(http.StatusOK, PresetResponse{Found: false, CamPresetsList: nil})
// 	}

// 	return ctx.JSON(http.StatusOK, PresetResponse{Found: true, Value: cam})
// }
