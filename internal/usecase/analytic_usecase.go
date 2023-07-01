package usecase

import (
	"context"
	"time"

	"github.com/krobus00/analytics-service/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type analyticUsecase struct {
	db           *gorm.DB
	analyticRepo model.AnalyticRepository
}

func NewAnalyticUsecase() model.AnalyticUsecase {
	return new(analyticUsecase)
}

func (uc *analyticUsecase) Create(ctx context.Context, payload *model.CreateAnalyticPayload) (*model.Analytic, error) {
	logger := logrus.WithFields(logrus.Fields{
		"source":   payload.Source,
		"medium":   payload.Medium,
		"campaign": payload.Campaign,
		"ip":       payload.IP,
		"country":  payload.Country,
		"city":     payload.City,
	})

	newAnalytic := &model.Analytic{
		Source:    payload.Source,
		Medium:    payload.Medium,
		Campaign:  payload.Campaign,
		IP:        payload.IP,
		Country:   payload.Country,
		City:      payload.City,
		CreatedAt: time.Now(),
	}

	err := uc.analyticRepo.Create(ctx, newAnalytic)
	if err != nil {
		logger.Error(err.Error())
		return nil, model.ErrCreateAnalytic
	}

	return newAnalytic, nil
}
