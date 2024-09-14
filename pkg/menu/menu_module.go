package menu

import (

	// "github.com/merger3/camserver/pkg/core"

	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/labstack/echo"
)

type MenuModule struct {
	Client  *twitch.Client
	Sources map[string]*Entry
	Cams    map[string]*CleanEntry
}

func NewMenuModule() *MenuModule {
	return &MenuModule{}
}

func (m MenuModule) RegisterRoutes(server *echo.Echo) {
	server.POST("/getSwapMenu", m.GetSwapMenu)
}

func (m *MenuModule) Init(resources map[string]any) {
	m.Sources = make(map[string]*Entry)
	m.Cams = make(map[string]*CleanEntry)
	m.Client = resources["twitch"].(*twitch.Client)

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
			case []interface{}:
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
			}
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