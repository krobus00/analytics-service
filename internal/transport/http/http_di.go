package http

import (
	"errors"

	"github.com/labstack/echo/v4"
)

func (t *HTTPDelivery) InjectEchoServer(e *echo.Echo) error {
	if e == nil {
		return errors.New("invalid echo")
	}
	t.httpServer = e
	return nil
}

func (t *HTTPDelivery) InjectAnalityCtrl(ctrl *AnalyticController) error {
	if ctrl == nil {
		return errors.New("invalid analytic controller")
	}
	t.analyticCtrl = ctrl
	return nil
}
