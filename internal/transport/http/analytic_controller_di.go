package http

import (
	"errors"

	"github.com/krobus00/analytics-service/internal/model"
)

func (t *AnalyticController) InjectAnalyticUC(uc model.AnalyticUsecase) error {
	if uc == nil {
		return errors.New("invalid analytic usecase")
	}
	t.analyticUC = uc
	return nil
}
