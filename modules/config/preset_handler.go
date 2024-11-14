package config

import (
	"fmt"
	"net/http"

	"github.com/merger3/camserver/modules/core"

	"github.com/labstack/echo/v4"
)

type PresetResponse struct {
	Found          bool        `json:"found"`
	CamPresetsList *CamPresets `json:"camPresets"`
}

func (c ConfigModule) CheckCacheSync(ctx echo.Context) error {
	if c.Cache.IsSynced {
		return ctx.JSON(http.StatusOK, map[string]any{
			"synced": true,
			"length": len(c.Cache.Cams),
		})
	} else {
		return ctx.JSON(http.StatusAccepted, map[string]any{
			"synced": true,
			"length": len(c.Cache.Cams),
		})
	}
}

func (c ConfigModule) GetCamPresets(ctx echo.Context) error {
	req := core.CamRequest{}

	if err := ctx.Bind(&req); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	if err := (&echo.DefaultBinder{}).BindHeaders(ctx, &req); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	for _, presets := range c.Cameras {
		if c.Aliases.ToBase(presets.CamName) == c.Aliases.ToBase(req.Cam) {
			return ctx.JSON(http.StatusOK, PresetResponse{Found: true, CamPresetsList: &presets})
		}
	}

	return ctx.JSON(http.StatusOK, PresetResponse{Found: false, CamPresetsList: nil})
}
