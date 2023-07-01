package http

import "github.com/labstack/echo/v4"

type HTTPDelivery struct {
	httpServer   *echo.Echo
	analyticCtrl *AnalyticController
}

func NewDelivery() *HTTPDelivery {
	return new(HTTPDelivery)
}

func (t *HTTPDelivery) InitRoutes() {
	t.httpServer.Use(parseHeaderMiddleware())
	api := t.httpServer.Group("/api")

	api.GET("", t.analyticCtrl.Create)
}
