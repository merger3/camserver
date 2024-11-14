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
	SwapMenu *CleanEntry `json:"swaps"`
}

func (m *MenuModule) GetSwapMenu(ctx echo.Context) error {
	req := core.CamRequest{}

	if err := ctx.Bind(&req); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	if err := (&echo.DefaultBinder{}).BindHeaders(ctx, &req); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	req.Cam = m.Aliases.ToCommon(m.Aliases.ToBase(m.Aliases.CleanName(req.Cam)))
	swaps, ok := m.Cams[req.Cam]
	if !ok {
		return ctx.JSON(http.StatusOK, SwapMenuResponse{Found: true, Cam: req.Cam, SwapMenu: m.Cams["pasture"]})
	}

	return ctx.JSON(http.StatusOK, SwapMenuResponse{Found: true, Cam: req.Cam, SwapMenu: swaps})
}
