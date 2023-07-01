package usecase

import (
	"errors"

	"github.com/krobus00/analytics-service/internal/model"
	"gorm.io/gorm"
)

func (uc *analyticUsecase) InjectDB(db *gorm.DB) error {
	if db == nil {
		return errors.New("invalid db")
	}
	uc.db = db
	return nil
}

func (uc *analyticUsecase) InjectAnalyticRepo(repo model.AnalyticRepository) error {
	if repo == nil {
		return errors.New("invalid repo")
	}
	uc.analyticRepo = repo
	return nil
}
