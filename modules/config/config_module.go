package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	// "github.com/merger3/camserver/modules/core"
	"github.com/merger3/camserver/managers/alias"
	"github.com/merger3/camserver/managers/cache"
	"github.com/merger3/camserver/managers/twitch"

	"github.com/labstack/echo/v4"
)

type CamPresets struct {
	CamName string `json:"name"`
	Presets []any  `json:"presets"`
}

type ButtonPreset struct {
	Name string `json:"name"`
}
type MenuPreset struct {
	Name       string       `json:"name"`
	Subentries []MenuPreset `json:"subentries"`
}
type HotkeyPreset struct {
	Name     string         `json:"name"`
	Hotkeys  string         `json:"hotkeys"`
	Sublayer []HotkeyPreset `json:"sublayer"`
}

type ConfigModule struct {
	Twitch        *twitch.TwitchManager
	Cache         *cache.CacheManager
	Aliases       alias.AliasManager
	ButtonPresets []CamPresets
	MenuPresets   []CamPresets
	HotkeyPresets []CamPresets
}

func NewConfigModule() *ConfigModule {
	return &ConfigModule{}
}

func (c ConfigModule) RegisterRoutes(server *echo.Echo) {
	server.POST("/api/camera/presets/buttons", c.GetButtonPresets)
	server.POST("/api/camera/presets/menus", c.GetMenuPresets)
	server.POST("/api/camera/presets/hotkeys", c.GetHotkeyPresets)
	server.POST("/api/authorize", c.GetAuthorized)
	server.GET("/api/synced", c.CheckCacheSync)
}

func (c *ConfigModule) Init(resources map[string]any) {
	c.ButtonPresets = LoadPresets("buttons")
	c.MenuPresets = LoadPresets("menus")
	c.HotkeyPresets = LoadPresets("hotkeys")
	c.Twitch = resources["twitch"].(*twitch.TwitchManager)
	c.Cache = resources["cache"].(*cache.CacheManager)
	c.Aliases = resources["aliases"].(alias.AliasManager)
}

func LoadPresets(set string) []CamPresets {
	file, err := os.Open(filepath.Join("configs", fmt.Sprintf("%s.presets.json", set)))
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	var p []CamPresets

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&p); err != nil {
		log.Fatalf("error unmarshalling JSON: %s", err)
	}

	return p
}
