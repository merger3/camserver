package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	// "github.com/merger3/camserver/modules/core"
	"github.com/merger3/camserver/managers/alias"
	"github.com/merger3/camserver/managers/cache"
	"github.com/merger3/camserver/managers/twitch"

	"github.com/labstack/echo/v4"
)

type Presets struct {
	Cameras []CamPresets `json:"cameras"`
}

type CamPresets struct {
	CamName string   `json:"name"`
	Presets []Preset `json:"presets"`
}

type Preset struct {
	Name       string   `json:"name"`
	Hotkeys    string   `json:"hotkeys"`
	Subentries []Preset `json:"subentries"`
	Sublayer   []Preset `json:"sublayer"`
}

type ConfigModule struct {
	Twitch  *twitch.TwitchManager
	Cache   *cache.CacheManager
	Aliases alias.AliasManager
	Presets
}

func NewConfigModule() *ConfigModule {
	return &ConfigModule{}
}

func (c ConfigModule) RegisterRoutes(server *echo.Echo) {
	server.POST("/api/camera/presets", c.GetCamPresets)
	server.POST("/api/authorize", c.GetAuthorized)
	server.GET("/api/synced", c.CheckCacheSync)
}

func (c *ConfigModule) Init(resources map[string]any) {
	c.Presets = LoadPresets()
	c.Twitch = resources["twitch"].(*twitch.TwitchManager)
	c.Cache = resources["cache"].(*cache.CacheManager)
	c.Aliases = resources["aliases"].(alias.AliasManager)
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
