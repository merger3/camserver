package menu

import (

	// "github.com/merger3/camserver/modules/core"

	"fmt"
	"net/http"

	"github.com/merger3/camserver/modules/core"

	"github.com/labstack/echo/v4"
)

type SwapMenuResponse struct {
	Found    bool        `json:"found"`
	Cam      string      `json:"cam"`
	Position int         `json:"position"`
	SwapMenu *CleanEntry `json:"swaps"`
}

func (m *MenuModule) GetSwapMenu(ctx echo.Context) error {
	req := core.Geom{}

	if err := ctx.Bind(&req); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	cam := m.Twitch.GetClickedCam(req)
	if !cam.Found {
		return ctx.JSON(http.StatusOK, SwapMenuResponse{})
	}

	swaps, ok := m.Cams[cam.Name]
	if !ok {
		return ctx.JSON(http.StatusOK, SwapMenuResponse{})
	}

	return ctx.JSON(http.StatusOK, SwapMenuResponse{Found: true, Cam: cam.Name, Position: cam.Position, SwapMenu: swaps})
}
