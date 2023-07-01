package http

import (
	"fmt"

	"github.com/krobus00/analytics-service/internal/constant"
	"github.com/krobus00/analytics-service/internal/model"
	"github.com/labstack/echo/v4"
)

type AnalyticController struct {
	analyticUC model.AnalyticUsecase
}

func NewAnalyticController() *AnalyticController {
	return new(AnalyticController)
}

func (t *AnalyticController) Create(eCtx echo.Context) error {
	var (
		ctx = buildContext(eCtx)
		res = model.NewDefaultResponse()
		req = new(model.HTTPCreateAnalyticPayload)
	)

	err := eCtx.Bind(req)
	if err != nil {
		res = model.WithBadRequestResponse(nil)
		return res.BuildResponse(eCtx)
	}

	_, err = t.analyticUC.Create(ctx, &model.CreateAnalyticPayload{
		Source:   req.Source,
		Medium:   req.Medium,
		Campaign: req.Campaign,
		IP:       fmt.Sprintf("%s", ctx.Value(constant.KeyUserIPCtx)),
		Country:  fmt.Sprintf("%s", ctx.Value(constant.KeyUserCountryCtx)),
		City:     fmt.Sprintf("%s", ctx.Value(constant.KeyUserCityCtx)),
	})

	if err != nil {
		return sanitizeError(err, eCtx)
	}

	return res.BuildResponse(eCtx)
}
