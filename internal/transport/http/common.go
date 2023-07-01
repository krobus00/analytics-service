package http

import (
	"context"
	"net/http"

	"github.com/krobus00/analytics-service/internal/constant"
	"github.com/krobus00/analytics-service/internal/model"
	"github.com/labstack/echo/v4"
)

func buildContext(eCtx echo.Context) context.Context {
	userIP := eCtx.Get(string(constant.KeyUserIPCtx))
	ctx := context.WithValue(eCtx.Request().Context(), constant.KeyUserIPCtx, userIP)

	userCountry := eCtx.Get(string(constant.KeyUserCountryCtx))
	ctx = context.WithValue(ctx, constant.KeyUserCountryCtx, userCountry)

	userCity := eCtx.Get(string(constant.KeyUserCityCtx))
	ctx = context.WithValue(ctx, constant.KeyUserCityCtx, userCity)
	return ctx
}

func sanitizeError(err error, eCtx echo.Context) error {
	switch v := err.(type) {
	case *model.Response:
		return eCtx.JSON(v.StatusCode, v)
	default:
		res := &model.Response{
			Message:    "internal server error",
			StatusCode: http.StatusInternalServerError,
		}
		return eCtx.JSON(res.StatusCode, res)
	}
}
