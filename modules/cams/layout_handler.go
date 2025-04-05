package cams

import (
	"fmt"
	"net/http"

	"github.com/merger3/camserver/modules/core"

	"github.com/labstack/echo/v4"
)

func (c CamsModule) GetLayout(ctx echo.Context) error {
	req := core.AuthHeaders{}

	if err := (&echo.DefaultBinder{}).BindHeaders(ctx, &req); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	// Make sure cache is as up to date as possible
	c.Twitch.Send(core.Command{User: req.User, Command: "!scenecams"})

	return ctx.JSON(http.StatusOK, c.Cache.Cams)
}
