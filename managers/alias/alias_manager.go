package alias

import (
	"fmt"
	"regexp"
	"strings"
)

type AliasManager struct {
	Aliases     map[string]string
	CommonNames map[string]string
}

func NewAliasManager() *AliasManager {
	return &AliasManager{Aliases: aliases, CommonNames: common}
}

func (a AliasManager) CleanName(input string) string {
	newInput := input
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Failed to condense input (%s): %v\n", input, r)
		}
	}()

	newInput = strings.ToLower(newInput)
	re1 := regexp.MustCompile(`e?s(\s|\W|$|multi(?:cam)?|cam|outdoor|indoor|inside|wideangle|corner|den)`)
	newInput = re1.ReplaceAllString(newInput, "$1")
	re2 := regexp.MustCompile(`(?:full)?cams?`)
	newInput = re2.ReplaceAllString(newInput, "")
	newInput = strings.ReplaceAll(newInput, " ", "")

	return newInput
}

func (a AliasManager) ToCommon(input string) string {
	common, ok := common[input]
	if !ok {
		return input
	} else {
		return common
	}
}

func (a AliasManager) ToBase(input string) string {
	base, ok := aliases[input]
	if !ok {
		return input
	} else {
		return base
	}
}

var common = map[string]string{
	"rat":            "rat1",
	"roach":          "roaches",
	"isopod":         "marty",
	"orangeisopod":   "bb",
	"crow":           "crowin",
	"crowoutdoor":    "crowout",
	"fox":            "foxes",
	"marmoset":       "marmout",
	"marmosetindoor": "marmin",
	"parrot":         "parrots",
}

var aliases = map[string]string{
	// pc
	"localpc":  "pc",
	"desktop":  "pc",
	"pclocal":  "pc",
	"alveuspc": "pc",
	"serverpc": "pc",
	"pc":       "pc",
	"pcserver": "pc",
	"remotepc": "pc",

	// phone
	"phone":       "phone",
	"alveusphone": "phone",
	"winnie":      "phone",
	"goat":        "phone",

	// puppy
	"puppy":    "puppy",
	"scorpion": "puppy",

	// roach
	"roach":   "roach",
	"barbara": "roach",

	// hank
	"hank":                         "hank",
	"mrmctrain":                    "hank",
	"choochoo":                     "hank",
	"hankthetankchoochoomrmctrain": "hank",

	// hankcorner
	"hankcorner": "hankcorner",
	"hank2":      "hankcorner",
	"mrmctrain2": "hankcorner",
	"choochoo2":  "hankcorner",
	"hanknight":  "hankcorner",

	// hankmulti
	"hankmulti":                     "hankmulti",
	"hankthetankchoochoomrmctrain3": "hankmulti",
	"hank3":                         "hankmulti",
	"mrmctrain3":                    "hankmulti",
	"choochoo3":                     "hankmulti",

	// isopod
	"isopod":      "isopod",
	"marty":       "isopod",
	"martyisopod": "isopod",

	// orangeisopod
	"orangeisopod":  "orangeisopod",
	"bb":            "orangeisopod",
	"bbisopod":      "orangeisopod",
	"sisopod":       "orangeisopod",
	"isopodorange":  "orangeisopod",
	"oisopod":       "orangeisopod",
	"spanishisopod": "orangeisopod",
	"isopod2":       "orangeisopod",

	// georgie
	"georgie": "georgie",
	"georg":   "georgie",

	// georgiewater
	"georgiewater":      "georgiewater",
	"georgieunderwater": "georgiewater",

	// nuthousebackup
	"nuthousebackup": "nuthousebackup",
	"nut":            "nuthousebackup",

	// servernuthouse
	"servernuthouse": "servernuthouse",
	"servernut":      "servernuthouse",
	"remotenut":      "servernuthouse",
	"remotenuthouse": "servernuthouse",

	// crow
	"crow":       "crow",
	"crowindoor": "crow",
	"crowin":     "crow",
	"crowinside": "crow",

	// crowoutdoor
	"crow2":       "crowoutdoor",
	"crowoutdoor": "crowoutdoor",
	"crowout":     "crowoutdoor",

	// crowmulti
	"crowmulti":          "crowmulti",
	"crow3":              "crowmulti",
	"crowoutcrowinmulti": "crowmulti",
	"crowoutcrowin":      "crowmulti",
	"crowoutcrow":        "crowmulti",
	"crowcrowin":         "crowmulti",

	// crowmulti2
	"crow4":              "crowmulti2",
	"crowmulti2":         "crowmulti2",
	"crowincrowoutmulti": "crowmulti2",
	"crowincrowout":      "crowmulti2",
	"crowcrowout":        "crowmulti2",
	"crowincrow":         "crowmulti2",

	// fox
	"fox": "fox",

	// foxmulti
	"fox2":              "foxmulti",
	"foxmulti":          "foxmulti",
	"foxfoxcorner":      "foxmulti",
	"foxfoxcornermulti": "foxmulti",

	// foxcorner
	"fox3":         "foxcorner",
	"foxcorner":    "foxcorner",
	"foxwideangle": "foxcorner",

	// foxden
	"fox4":   "foxden",
	"foxden": "foxden",

	// marmoset
	"marmoset":        "marmoset",
	"marmosetoutdoor": "marmoset",
	"marm":            "marmoset",
	"marmoutdoor":     "marmoset",
	"marmsout":        "marmoset",
	"marmout":         "marmoset",

	// marmosetindoor
	"marmoset2":      "marmosetindoor",
	"marmosetindoor": "marmosetindoor",
	"marmindoor":     "marmosetindoor",
	"marminside":     "marmosetindoor",
	"marmsin":        "marmosetindoor",
	"marmin":         "marmosetindoor",

	// marmosetmulti
	"marmoset3":                  "marmosetmulti",
	"marmosetmulti":              "marmosetmulti",
	"marmmulti":                  "marmosetmulti",
	"marmoutmarmin":              "marmosetmulti",
	"marmmarmin":                 "marmosetmulti",
	"marmoutmarm":                "marmosetmulti",
	"marmoutmarminmulti":         "marmosetmulti",
	"marmosetoutmarmosetinmulti": "marmosetmulti",

	// scenes
	"4outdoor":     "pasture parrot marmoset fox",
	"4outdoorcam":  "pasture parrot marmoset fox",
	"multioutdoor": "pasture parrot marmoset fox",
	"night":        "wolf pasture parrot fox crow marmoset",
	"outdoor":      "wolf pasture parrot fox crow marmoset",
	"outside":      "wolf pasture parrot fox crow marmoset",
	"live":         "wolf pasture parrot fox crow marmoset",
	"nightbig":     "wolf pasture parrot fox crow marmoset",
	"outdoorbig":   "wolf pasture parrot fox crow marmoset",
	"outsidebig":   "wolf pasture parrot fox crow marmoset",
	"nightb":       "wolf pasture parrot fox crow marmoset",
	"night2":       "wolf pasture parrot fox crow marmoset",
	"ncb":          "wolf pasture parrot fox crow marmoset",
	"indoor":       "georgie hank puppy chin isopod roach",
	"4cam":         "georgie hank puppy chin isopod roach",
	"inside":       "georgie hank puppy chin isopod roach",
	"indoorbig":    "georgie hank puppy chin isopod roach",
	"insidebig":    "georgie hank puppy chin isopod roach",

	// chin
	"chin":       "chin",
	"chinchilla": "chin",
	"snork":      "chin",
	"moomin":     "chin",
	"fluffy":     "chin",

	// rat
	"rat":    "rat",
	"rat1":   "rat",
	"rattop": "rat",
	"nilla":  "rat",
	"chip":   "rat",
	"ratt":   "rat",

	// rat2
	"rat2":      "rat2",
	"ratmiddle": "rat2",
	"ratm":      "rat2",

	// rat3
	"rat3":      "rat3",
	"ratbottom": "rat3",
	"ratb":      "rat3",

	// ratmulti
	"ratmulti": "ratmulti",
	"rat4":     "ratmulti",
	"ratall":   "rat rat2 rat3",
	"ratstack": "rat rat2 rat3",

	// connorpc
	"connorpc":      "connorpc",
	"connordesktop": "connorpc",

	// construction
	"construction": "construction",
	"timelapse":    "construction",

	// connorintro
	"connorintro": "connorintro",
	"peni":        "connorintro",

	// accintro
	"accintro": "accintro",
	"acintro":  "accintro",

	// accbrb
	"accbrb": "accbrb",
	"acbrb":  "accbrb",

	// accending
	"accending": "accending",
	"acending":  "accending",
	"accend":    "accending",
	"acend":     "accending",

	// ccintro
	"ccintro": "ccintro",
	"cintro":  "ccintro",

	// ccbrb
	"ccbrb": "ccbrb",
	"cbrb":  "ccbrb",

	// ccending
	"ccending": "ccending",
	"cending":  "ccending",
	"ccend":    "ccending",
	"cend":     "ccending",

	// nickending
	"nickending": "nickending",
	"nickend":    "nickending",

	// chatchat
	"chatchat":    "chatchat",
	"bugmic":      "chatchat",
	"chatchatmic": "chatchat",

	// phonemic
	"phonemic":   "phonemic",
	"phoneaudio": "phonemic",
	"mobilemic":  "phonemic",

	// wolf
	"wolf":        "wolf",
	"wolv":        "wolf",
	"timber":      "wolf",
	"awa":         "wolf",
	"wolfo":       "wolf",
	"wolfout":     "wolf",
	"wolfoutdoor": "wolf",
	"wolfoutside": "wolf",

	// wolfcorner
	"wolfcorner": "wolfcorner",
	"wolf2":      "wolfcorner",
	"wolvcorner": "wolfcorner",
	"wolfside":   "wolfcorner",
	"wolfdeck":   "wolfcorner",

	// wolfden
	"wolfden":     "wolfden",
	"wolf3":       "wolfden",
	"wolvden":     "wolfden",
	"wolfpondden": "wolfden",

	// wolfden2
	"wolfden2":    "wolfden2",
	"wolf4":       "wolfden2",
	"wolvden2":    "wolfden2",
	"wolfdeckden": "wolfden2",

	// wolfindoor
	"wolvindoor": "wolfindoor",
	"wolf5":      "wolfindoor",
	"wolfindoor": "wolfindoor",
	"wolfinside": "wolfindoor",
	"wolfin":     "wolfindoor",
	"wolfi":      "wolfindoor",

	// wolfmulti
	"wolfmulti":           "wolfmulti",
	"wolf6":               "wolfmulti",
	"wolvmulti":           "wolfmulti",
	"wolfoutmulti":        "wolfmulti",
	"wolfwolfcornermulti": "wolfmulti",

	// wolfmulti2
	"wolf7":           "wolfmulti2",
	"wolfindoormulti": "wolfmulti2",
	"wolfinmulti":     "wolfmulti2",
	"wolfinsidemulti": "wolfmulti2",
	"wolfwolfinmulti": "wolfmulti2",
	"wolfwolfin":      "wolfmulti2",

	// wolfmulti3
	"wolf8":            "wolfmulti3",
	"wolfdenmulti":     "wolfmulti3",
	"wolfwolfdenmulti": "wolfmulti3",
	"wolfwolfden":      "wolfmulti3",

	// wolfmulti4
	"wolf9":             "wolfmulti4",
	"wolfden2multi":     "wolfmulti4",
	"wolfwolfden2multi": "wolfmulti4",
	"wolfwolfden2":      "wolfmulti4",

	// wolfmulti5
	"wolf10":                "wolfmulti5",
	"wolfcornermulti":       "wolfmulti5",
	"wolfcornerwolfin":      "wolfmulti5",
	"wolfcornerwolfinmulti": "wolfmulti5",
	"wolfcwolfin":           "wolfmulti5",
	"wolfcwolfi":            "wolfmulti5",
}
