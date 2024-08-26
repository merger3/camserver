package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	// "github.com/merger3/camserver/pkg/core"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/labstack/echo"
)

type ConfigModule struct {
	Client *twitch.Client
	Presets
}

type Presets struct {
	Cameras []CamPresets `json:"cameras"`
}

type CamPresets struct {
	CamName string   `json:"name"`
	Presets []string `json:"presets"`
}

func NewConfigModule() *ConfigModule {
	return &ConfigModule{}
}

func (c ConfigModule) RegisterRoutes(server *echo.Echo) {
	server.POST("/getConfig", c.GetCamPresets)
	server.POST("/getClickedCam", c.GetClickedCamPresets)
}

func (c *ConfigModule) Init(resources map[string]any) {
	c.Presets = LoadPresets()
	c.Client = resources["twitch"].(*twitch.Client)
}

func LoadPresets() Presets {
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

	return p
}
