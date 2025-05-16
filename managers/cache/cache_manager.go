package cache

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/merger3/camserver/managers/alias"
	"github.com/merger3/camserver/modules/core"
)

type CacheManager struct {
	Cams              []string
	Aliases           alias.AliasManager
	LastSynced        time.Time
	LastAttemptedSync time.Time
	SyncAttempts      float64
	IsSynced          bool
	HTTPClient        *http.Client
	APIKey            string
}

func NewCacheManager(client *http.Client, token string) *CacheManager {
	newCM := &CacheManager{Cams: make([]string, 6), HTTPClient: client, APIKey: token}
	// newCM.SyncCache()
	// go newCM.HydrateCache()
	return newCM
}

func (cm *CacheManager) HydrateCache() error {
	ticker := time.NewTicker(5 * time.Minute)
	for range ticker.C {
		cm.SyncCache()
	}
	return nil
}

func (cm *CacheManager) SyncCache() error {
	url := "https://api.ptz.app:2053/api/command"

	// str := "Scene: Custom Cams: 1: wolf, 2: pasture, 3: wolfmulti, 4: pushpop, 5: pushpopindoor, 6: crowin"
	// cm.ParseScene(str)
	// return nil

	requestBody, err := json.Marshal(core.Payload{Message: "!scenecams"})
	if err != nil {
		return err
	}

	request, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cm.APIKey))
	request.Header.Set("Content-Type", "application/json")

	rsp, err := cm.HTTPClient.Do(request)
	if err != nil {
		return err
	}

	defer rsp.Body.Close()

	if rsp.StatusCode == http.StatusOK {
		// Parse the response body
		var response core.Payload
		if err := json.NewDecoder(rsp.Body).Decode(&response); err != nil {
			return err
		}

		// cm.ParseScene("Scene: customcamsbig Cams: 1: crowoutdoor, 2: pushpop, 3: marmosetindoor, 4: pasture, 5: wolfswitch, 6: wolf")
		cm.ParseScene(strings.ReplaceAll(response.Message, "\n", ""))
	} else {
		return core.ErrFailedToSyncCacheWithAPI
	}

	return nil

}

func isNumber(s string) bool {
	var d = regexp.MustCompile(`^[0-9]+$`)
	return d.MatchString(s)
}

func (cm *CacheManager) ParseScene(scenecamsRaw string) {
	fmt.Println("parsing")
	scenecamsRE := regexp.MustCompile(`^Scene: \w+ Current Scene: \w+ Cams: ((\d: \w+,? ?)+)$`)
	if !scenecamsRE.MatchString(scenecamsRaw) {
		return
	}
	fmt.Println("matched regex")
	matches := scenecamsRE.FindStringSubmatch(scenecamsRaw)
	scenecams := matches[1]

	camsArray := strings.Split(strings.ReplaceAll(scenecams, " ", ""), ",")
	if len(camsArray) < 6 {
		fmt.Println("array less than 6")
		return
	}

	newArray := make([]string, 0)
	for i := 0; i < len(camsArray); i++ {
		if len(camsArray[i]) > 2 {
			newArray = append(newArray, camsArray[i][2:])
		}
	}
	cm.Cams = newArray
	cm.IsSynced = true
	cm.LastSynced = time.Now()
	fmt.Println(cm.Cams)
}

func (cm CacheManager) ProcessSwap(first, second string) {
	switch {
	case isNumber(first) && isNumber(second):
		cm.swapByIndex(first, second)
	case !isNumber(first) && isNumber(second):
		first = cm.Aliases.ToBase(cm.Aliases.CleanName(first))
		secondInt, _ := strconv.Atoi(second)
		cm.swapByNameAndIndex(secondInt, first)
	case isNumber(first) && !isNumber(second):
		firstInt, _ := strconv.Atoi(first)
		second = cm.Aliases.ToBase(cm.Aliases.CleanName(second))
		cm.swapByNameAndIndex(firstInt, second)
	case !isNumber(first) && !isNumber(second):
		first = cm.Aliases.ToBase(cm.Aliases.CleanName(first))
		second = cm.Aliases.ToBase(cm.Aliases.CleanName(second))
		cm.swapByName(first, second)
	}
}

func (cm CacheManager) swapByName(first, second string) {
	iFirst := slices.Index(cm.Cams, first)
	iSecond := slices.Index(cm.Cams, second)

	switch {
	case iFirst != -1 && iSecond == -1:
		cm.Cams[iFirst] = second
	case iFirst == -1 && iSecond != -1:
		cm.Cams[iSecond] = first
	case iFirst != -1 && iSecond != -1:
		cm.Cams[iFirst] = second
		cm.Cams[iSecond] = first
	}
}

func (cm CacheManager) swapByIndex(first, second string) {
	// These can't fail, checking already occurred
	iFirst, _ := strconv.Atoi(first)
	iSecond, _ := strconv.Atoi(second)
	if (iFirst < 1) || (iFirst > 6) {
		return
	}
	if (iSecond < 1) || (iSecond > 6) {
		return
	}

	iFirst--
	iSecond--

	tmp := cm.Cams[iFirst]
	cm.Cams[iFirst] = cm.Cams[iSecond]
	cm.Cams[iSecond] = tmp
}

func (cm CacheManager) swapByNameAndIndex(first int, second string) {
	if (first < 1) || (first > 6) {
		return
	}
	first = first - 1
	iSecond := slices.Index(cm.Cams, second)

	if iSecond == -1 {
		cm.Cams[first] = second
	} else {
		tmp := cm.Cams[first]
		cm.Cams[first] = cm.Cams[iSecond]
		cm.Cams[iSecond] = tmp
	}
}

func (cm *CacheManager) Invalidate() {
	cm.IsSynced = false
}

func (cm CacheManager) FetchFromCache(position int) core.ClickedCam {
	if position < 1 || position > len(cm.Cams) {
		return core.ClickedCam{}
	} else {
		return core.ClickedCam{Found: true, Name: cm.Cams[position-1], Position: position}
	}
}
