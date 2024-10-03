package config

import (
	"fmt"
	"net/http"

	"github.com/merger3/camserver/modules/core"

	"github.com/labstack/echo/v4"
)

type AuthResponse struct {
	Authorized bool `json:"authorized"`
}

func (c ConfigModule) GetAuthorized(ctx echo.Context) error {
	req := core.AuthHeaders{}

	if err := (&echo.DefaultBinder{}).BindHeaders(ctx, &req); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	authMap := createAuthMap()
	fmt.Println(req.User)
	fmt.Println(authMap[req.User])
	return ctx.JSON(http.StatusOK, AuthResponse{Authorized: authMap[req.User]})
}

func createAuthMap() map[string]bool {

	// Define user lists
	commandAdmins := []string{"spacevoyage", "maya", "theconnorobrien", "alveussanctuary"}
	commandSuperUsers := []string{"ellaandalex", "dionysus1911", "dannydv", "maxzillajr", "illjx", "kayla_alveus",
		"alex_b_patrick", "lindsay_alveus", "strickknine", "tarantulizer", "spiderdaynightlive",
		"srutiloops", "evantomology", "amanda2815"}
	commandMods := []string{"pjeweb", "loganrx_", "MattIPv4", "Mik_MWP"}
	commandOperator := []string{"96allskills", "stolenarmy_", "berlac", "dansza", "loganrx_", "merger3", "nitelitedf",
		"purplemartinconservation", "wazix11", "lazygoosepxls", "alxiszzz", "shutupleonard",
		"taizun", "lumberaxe1", "glennvde", "wolfone_", "dohregard", "lakel1", "darkrow_",
		"minipurrl", "gnomechildboi", "danman149", "hunnybeehelen", "strangecyan"}
	commandVips := []string{"tfries_", "sivvii_", "ghandii_", "axialmars", "jazz_peru", "stealfydoge",
		"xano218", "experimentalcyborg", "klav___", "monkarooo", "nixxform", "madcharliekelly",
		"josh_raiden", "jateu", "storesE6", "rebecca_h9", "matthewde", "user_11_11", "huniebeexd",
		"kurtyykins", "breacherman", "bryceisrightjr", "sumaxu", "mariemellie", "ewok_626",
		"quokka64", "casualruffian", "likethecheesebri", "viphippo", "bagel_deficient", "otsargh",
		"just_some_donkus", "fiveacross", "itszalndrin", "nicoleeverleigh", "fishymeep", "ponchobee"}

	// Create the map and add all users from the lists
	userMap := make(map[string]bool)

	for _, admin := range commandAdmins {
		userMap[admin] = true
	}

	for _, superUser := range commandSuperUsers {
		userMap[superUser] = true
	}

	for _, mod := range commandMods {
		userMap[mod] = true
	}

	for _, operator := range commandOperator {
		userMap[operator] = true
	}

	for _, vip := range commandVips {
		userMap[vip] = true
	}

	// Output the map
	return userMap
}
