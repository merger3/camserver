package config

import (
	"fmt"
	"net/http"

	"github.com/merger3/camserver/modules/click"
	"github.com/merger3/camserver/modules/core"

	"github.com/labstack/echo/v4"
)

type PresetRequest struct {
	Cam string `json:"camera"`
}

type PresetResponse struct {
	Found          bool        `json:"found"`
	CamPresetsList *CamPresets `json:"camPresets"`
}

func (c ConfigModule) GetCamPresets(ctx echo.Context) error {
	req := PresetRequest{}

	if err := ctx.Bind(&req); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	for _, presets := range c.Cameras {
		if presets.CamName == req.Cam {
			return ctx.JSON(http.StatusOK, PresetResponse{Found: true, CamPresetsList: &presets})
		}
	}

	return ctx.JSON(http.StatusOK, PresetResponse{Found: false, CamPresetsList: nil})
}

func (c ConfigModule) GetClickedCamPresets(ctx echo.Context) error {
	req := core.Geom{}

	if err := ctx.Bind(&req); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	cam := click.GetClickedCam(c.Client, req)
	if !cam.Found {
		return ctx.JSON(http.StatusOK, PresetResponse{Found: false, CamPresetsList: nil})
	}

	for _, presets := range c.Cameras {
		if presets.CamName == cam.Name {
			return ctx.JSON(http.StatusOK, PresetResponse{Found: true, CamPresetsList: &presets})
		}
	}

	return ctx.JSON(http.StatusOK, PresetResponse{Found: false, CamPresetsList: nil})
}