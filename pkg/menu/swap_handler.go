package menu

import (

	// "github.com/merger3/camserver/pkg/core"

	"fmt"
	"net/http"

	"github.com/merger3/camserver/pkg/click"
	"github.com/merger3/camserver/pkg/core"

	"github.com/labstack/echo"
)

type SwapMenuResponse struct {
	Found    bool        `json:"found"`
	Cam      string      `json:"cam"`
	SwapMenu *CleanEntry `json:"swaps"`
}

func (m *MenuModule) GetSwapMenu(ctx echo.Context) error {
	req := core.Geom{}

	if err := ctx.Bind(&req); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	cam := click.GetClickedCam(m.Client, req)
	// cam := "parrots"
	if cam == "" {
		return ctx.JSON(http.StatusOK, SwapMenuResponse{Found: false, Cam: cam, SwapMenu: nil})
	}

	swaps, ok := m.Cams[cam]
	if !ok {
		return ctx.JSON(http.StatusOK, SwapMenuResponse{Found: false, Cam: cam, SwapMenu: nil})
	}

	return ctx.JSON(http.StatusOK, SwapMenuResponse{Found: true, Cam: cam, SwapMenu: swaps})
}
