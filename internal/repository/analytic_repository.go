package repository

import (
	"context"

	"github.com/krobus00/analytics-service/internal/model"
	"github.com/krobus00/analytics-service/internal/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type analyticRepository struct {
	db *gorm.DB
}

func NewAnalyticRepository() model.AnalyticRepository {
	return new(analyticRepository)
}

func (r *analyticRepository) Create(ctx context.Context, analytic *model.Analytic) error {
	if analytic.ID == "" {
		analytic.ID = utils.GenerateUUID()
	}

	logger := logrus.WithFields(logrus.Fields{
		"id":       analytic.ID,
		"source":   analytic.Source,
		"medium":   analytic.Medium,
		"campaign": analytic.Campaign,
		"ip":       analytic.IP,
		"country":  analytic.Country,
		"city":     analytic.City,
	})

	err := r.db.WithContext(ctx).Create(&analytic).Error
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
