//go:generate mockgen -destination=mock/mock_analytic_repository.go -package=mock github.com/krobus00/analytics-service/internal/model ProductRepository
//go:generate mockgen -destination=mock/mock_analytic_usecase.go -package=mock github.com/krobus00/analytics-service/internal/model ProductUsecase

package model

import (
	"context"
	"net/http"
	"time"

	"gorm.io/gorm"
)

var (
	ErrCreateAnalytic = &Response{
		Message:    "error when create new analytic data",
		StatusCode: http.StatusInternalServerError,
	}
)

type Analytic struct {
	ID        string `gorm:"primaryKey"`
	Source    string
	Medium    string
	Campaign  string
	IP        string
	Country   string
	City      string
	CreatedAt time.Time `gorm:"<-:create"` // read and create
}

func (Analytic) TableName() string {
	return "analytics"
}

type CreateAnalyticPayload struct {
	Source   string
	Medium   string
	Campaign string
	IP       string
	Country  string
	City     string
}

type HTTPCreateAnalyticPayload struct {
	Source   string `query:"utm_source"`
	Medium   string `query:"utm_medium"`
	Campaign string `query:"utm_campaign"`
}

type AnalyticRepository interface {
	Create(ctx context.Context, analytic *Analytic) error

	// DI
	InjectDB(db *gorm.DB) error
}

type AnalyticUsecase interface {
	Create(ctx context.Context, payload *CreateAnalyticPayload) (*Analytic, error)

	// DI
	InjectDB(db *gorm.DB) error
	InjectAnalyticRepo(repo AnalyticRepository) error
}
