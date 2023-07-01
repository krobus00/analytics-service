package http

import (
	"github.com/krobus00/analytics-service/internal/constant"
	"github.com/labstack/echo/v4"
)

func parseHeaderMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(eCtx echo.Context) error {
			req := eCtx.Request()
			headers := req.Header

			eCtx.Set(string(constant.KeyUserIPCtx), headers.Get("CF-Connecting-IP"))
			eCtx.Set(string(constant.KeyUserCountryCtx), headers.Get("CF-IPCountry"))
			eCtx.Set(string(constant.KeyUserCityCtx), headers.Get("CF-IPCity"))
			return next(eCtx)
		}
	}
}
