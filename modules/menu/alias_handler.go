package menu

import (

	// "github.com/merger3/camserver/modules/core"

	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AliasRequest struct {
	Cam string `json:"cam"`
}

func (m *MenuModule) GetAlias(ctx echo.Context) error {
	req := AliasRequest{}

	if err := ctx.Bind(&req); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"result": m.Aliases.ToBase(m.Aliases.CleanName(req.Cam)),
	})
}
