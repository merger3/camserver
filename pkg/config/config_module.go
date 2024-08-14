package config

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo"
)

type ConfigManager struct {
	Presets
}

type Presets struct {
	Cameras []CamPresets `json:"cameras"`
}

type CamPresets struct {
	CamName string   `json:"camName"`
	Presets []string `json:"presets"`
}

type PresetRequest struct {
	Cam string `json:"camera"`
}

type PresetResponse struct {
	Found          bool        `json:"found"`
	CamPresetsList *CamPresets `json:"camPresets"`
}

func NewConfigManager() ConfigManager {
	file, err := os.Open(filepath.Join("configs", "presets.json"))
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	var p Presets

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&p); err != nil {
		log.Fatalf("error unmarshalling JSON: %s", err)
	}

	return ConfigManager{p}
}

func (c ConfigManager) GetPresets(ctx echo.Context) error {
	req := PresetRequest{}

	if err := ctx.Bind(&req); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	file, err := os.Open(filepath.Join("configs", "presets.json"))
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	var p Presets

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&p); err != nil {
		log.Fatalf("error unmarshalling JSON: %s", err)
	}

	// This is slow, but Cameras should never be long enough for it to matter
	for _, presets := range p.Cameras {
		if presets.CamName == req.Cam {
			return ctx.JSON(http.StatusOK, PresetResponse{Found: true, CamPresetsList: &presets})
		}
	}

	return ctx.JSON(http.StatusOK, PresetResponse{Found: false, CamPresetsList: nil})
}
