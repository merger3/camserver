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

	fmt.Println(req.User)
	fmt.Println(c.Twitch.AuthMap[req.User])
	return ctx.JSON(http.StatusOK, AuthResponse{Authorized: c.Twitch.AuthMap[req.User]})
}
