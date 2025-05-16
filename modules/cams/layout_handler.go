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

	fmt.Printf("pre conversion: %+v\n", c.Cache.Cams)
	layout := c.convertLayoutNames(c.Cache.Cams)
	fmt.Printf("post conversion: %+v\n", layout)
	// Make sure cache is as up to date as possible
	// c.Twitch.Send(core.Command{User: req.User, Command: "!scenecams"})

	return ctx.JSON(http.StatusOK, c.Cache.Cams)
}

func (c CamsModule) convertLayoutNames(layout []string) []string {
	newLayout := make([]string, len(layout))
	for i, name := range layout {
		newLayout[i] = c.Aliases.ToBase(name)
	}
	return newLayout
}
