package menu

import (

	// "github.com/merger3/camserver/modules/core"

	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"

	"github.com/labstack/echo/v4"
	"github.com/merger3/camserver/managers/alias"
	"github.com/merger3/camserver/managers/twitch"
)

type MenuModule struct {
	Twitch  *twitch.TwitchManager
	Sources map[string]*Entry
	Cams    map[string]*CleanEntry
	Aliases alias.AliasManager
}

func NewMenuModule() *MenuModule {
	return &MenuModule{}
}

func (m MenuModule) RegisterRoutes(server *echo.Echo) {
	server.POST("/api/camera/swaps", m.GetSwapMenu)
}

func (m *MenuModule) Init(resources map[string]any) {
	m.Sources = make(map[string]*Entry)
	m.Cams = make(map[string]*CleanEntry)
	m.Twitch = resources["twitch"].(*twitch.TwitchManager)
	m.Aliases = resources["aliases"].(alias.AliasManager)

	m.LoadSource("base")
	PopulateEntries(m.Sources["base"], m.Sources["base"])

	m.LoadSource("cams")
	for _, e := range m.Sources["cams"].SubEntries {
		e := &e
		PopulateEntries(m.Sources["base"], e)
		m.ApplyMods(e)

		ce := &CleanEntry{}
		ce.CopyFromEntry(e)
		m.Cams[e.Label] = ce
	}

}

type ModAction string

const (
	Add     ModAction = "add"
	Remove  ModAction = "remove"
	Move    ModAction = "move"
	Rename  ModAction = "rename"
	Flatten ModAction = "flatten"
	Swap    ModAction = "swap"
)

type Entry struct {
	Label      string   `json:"label"`
	Import     []string `json:"import"`
	SubEntries []Entry  `json:"subentries"`
	ModList    []Mod    `json:"mods"`
}

type Mod struct {
	Action ModAction      `json:"action"`
	Target []string       `json:"target"`
	Props  map[string]any `json:"props"`
}

type CleanEntry struct {
	Label      string       `json:"label"`
	SubEntries []CleanEntry `json:"subentries"`
}

func (m *MenuModule) LoadSource(name string) {
	file, err := os.Open(filepath.Join("configs", fmt.Sprintf("%s.swaps.json", name)))
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	m.Sources[name] = &Entry{Label: name}
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&m.Sources[name].SubEntries); err != nil {
		log.Fatalf("error unmarshalling JSON: %s", err)
	}
}

func (dst *CleanEntry) CopyFromEntry(src *Entry) {
	dst.Label = src.Label
	if len(src.SubEntries) != 0 {
		dst.SubEntries = make([]CleanEntry, len(src.SubEntries))
		for i, v := range src.SubEntries {
			dst.SubEntries[i].CopyFromEntry(&v)
		}
	}
}

func convertToStringSlice(input interface{}) ([]string, error) {
	itemInterface, ok := input.([]interface{})
	if !ok {
		return nil, errors.New("expected []interface{} for input")
	}

	itemStrings := make([]string, len(itemInterface))
	for i, v := range itemInterface {
		str, ok := v.(string)
		if !ok {
			return nil, errors.New("expected string element in input slice")
		}
		itemStrings[i] = str
	}

	return itemStrings, nil
}

func CopyImports(dst *Entry, src Entry) {
	dst.SubEntries = make([]Entry, len(src.SubEntries))
	copy(dst.SubEntries, src.SubEntries)
	if dst.Label == "" {
		dst.Label = src.Label
	}
	for i, v := range src.SubEntries {
		if len(v.SubEntries) != 0 {
			CopyImports(&dst.SubEntries[i], v)
		}
	}
}

func PopulateEntries(source, entry *Entry) {
	for i := range entry.SubEntries {
		subEntry := &entry.SubEntries[i]
		if len(subEntry.SubEntries) != 0 {
			PopulateEntries(source, subEntry)
		}
		if len(subEntry.Import) != 0 {
			targetIndex, targetParent := findTarget(source, subEntry.Import...)
			if targetIndex == -1 {
				continue
			}
			importedEntry := targetParent.SubEntries[targetIndex]
			CopyImports(subEntry, importedEntry)

		}
	}
}

func findTarget(entry *Entry, targetPath ...string) (int, *Entry) {
	var idx int
	for i, step := range targetPath {
		idx = slices.IndexFunc(entry.SubEntries, func(c Entry) bool { return c.Label == step })
		if idx == -1 {
			return -1, nil
		} else if i != len(targetPath)-1 {
			entry = &entry.SubEntries[idx]
		}
	}
	return idx, entry
}

func (m MenuModule) ApplyMods(entry *Entry) {
	for _, mod := range entry.ModList {
		targetIndex, targetParent := findTarget(entry, mod.Target...)
		if targetIndex == -1 {
			continue
		}
		target := &targetParent.SubEntries[targetIndex]
		switch mod.Action {
		case Add:
			source, ok := m.Sources[mod.Props["source"].(string)]
			if !ok {
				continue
			}
			itemStrings, err := convertToStringSlice(mod.Props["item"])
			if err != nil {
				panic(err)
			}
			newEntryIndex, newEntryParent := findTarget(source, itemStrings...)
			if newEntryIndex == -1 {
				continue
			}
			position := 0
			if targetIndex != 0 {
				position = targetIndex + 1
			}
			newEntrySource := newEntryParent.SubEntries[newEntryIndex]
			targetParent.SubEntries = slices.Insert(targetParent.SubEntries, position, newEntrySource)
		case Remove:
			targetParent.SubEntries = slices.Delete(targetParent.SubEntries, targetIndex, targetIndex+1)
			for len(mod.Target) > 1 && len(targetParent.SubEntries) == 0 {
				mod.Target = mod.Target[:len(mod.Target)-1]
				targetIndex, targetParent = findTarget(entry, mod.Target...)
				if targetIndex == -1 {
					continue
				}
				targetParent.SubEntries = slices.Delete(targetParent.SubEntries, targetIndex, targetIndex+1)
			}
		case Move:
			switch mod.Props["position"].(type) {
			case float64:
				tmpTarget := *target
				position := int(mod.Props["position"].(float64))
				targetParent.SubEntries = slices.Delete(targetParent.SubEntries, targetIndex, targetIndex+1)
				if position > len(targetParent.SubEntries) {
					position = len(targetParent.SubEntries)
				}
				targetParent.SubEntries = slices.Insert(targetParent.SubEntries, position, tmpTarget)
			case string:
				var position int
				switch mod.Props["position"].(string) {
				case "top", "start", "begin", "beginning":
					position = 0
				case "bottom", "end", "ending":
					position = len(targetParent.SubEntries)
				case "middle", "center":
					position = len(targetParent.SubEntries) / 2
				default:
					continue
				}
				tmpTarget := *target
				targetParent.SubEntries = slices.Delete(targetParent.SubEntries, targetIndex, targetIndex+1)
				if position > len(targetParent.SubEntries) {
					position = len(targetParent.SubEntries)
				}
				targetParent.SubEntries = slices.Insert(targetParent.SubEntries, position, tmpTarget)
			case []interface{}:
				itemStrings, err := convertToStringSlice(mod.Props["position"])
				if err != nil {
					panic(err)
				}
				swapEntryIndex, swapEntryParent := findTarget(entry, itemStrings...)
				if swapEntryIndex == -1 {
					fmt.Println("continuing...")
					continue
				}
				swapEntry := &swapEntryParent.SubEntries[swapEntryIndex]

				location := len(swapEntry.SubEntries)
				switch mod.Props["location"].(type) {
				case float64:
					location = int(mod.Props["location"].(float64))
					if location < 0 {
						location = 0
					} else if location > len(targetParent.SubEntries) {
						location = len(targetParent.SubEntries)
					}
				case string:
					switch mod.Props["location"].(string) {
					case "top", "start", "begin", "beginning":
						location = 0
					case "bottom", "end", "ending":
						location = len(swapEntry.SubEntries)
					case "middle", "center":
						location = len(swapEntry.SubEntries) / 2
					}
				case map[string]interface{}:
					anchor, ok := mod.Props["location"].(map[string]any)["anchor"].(string)
					if !ok {
						fmt.Println("Failed to get anchor")
						break
					}
					relation, ok := mod.Props["location"].(map[string]any)["relation"].(string)
					if !ok {
						fmt.Println("Failed to get relation")
						break
					}
					anchorEntryIndex, _ := findTarget(swapEntry, anchor)
					if anchorEntryIndex == -1 {
						fmt.Println("Unable to find location")
						break
					}

					switch relation {
					case "above":
						location = anchorEntryIndex
					case "below":
						location = anchorEntryIndex + 1
					}

				}

				swapEntry.SubEntries = slices.Insert(swapEntry.SubEntries, location, *target)
				targetParent.SubEntries = slices.Delete(targetParent.SubEntries, targetIndex, targetIndex+1)
			}

		case Swap:
			itemStrings, err := convertToStringSlice(mod.Props["position"])
			if err != nil {
				panic(err)
			}
			swapEntryIndex, swapEntryParent := findTarget(entry, itemStrings...)
			if swapEntryIndex == -1 {
				continue
			}
			swapEntry := swapEntryParent.SubEntries[swapEntryIndex]
			tmpTarget := *target
			targetParent.SubEntries[targetIndex] = swapEntry
			swapEntryParent.SubEntries[swapEntryIndex] = tmpTarget
		case Rename:
			target.Label = mod.Props["name"].(string)
		case Flatten:
			targetParent.SubEntries = slices.Insert(targetParent.SubEntries, targetIndex, target.SubEntries...)
			targetIndex, targetParent = findTarget(entry, mod.Target...)
			targetParent.SubEntries = slices.Delete(targetParent.SubEntries, targetIndex, targetIndex+1)
		default:
			continue
		}
	}
}
