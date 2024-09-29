package config

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CacheRequest struct {
	Position int `json:"position"`
}

type CacheResponse struct {
	Found bool   `json:"found"`
	Cam   string `json:"camera"`
}

func (c ConfigModule) CheckCache(ctx echo.Context) error {
	req := CacheRequest{}

	if err := ctx.Bind(&req); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	resp := CacheResponse{Cam: c.Cache.FetchFromCache(req.Position)}

	if resp.Cam == "" {
		resp.Found = false
	} else {
		resp.Found = true
	}

	return ctx.JSON(http.StatusOK, resp)
}
