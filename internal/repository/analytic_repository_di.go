package repository

import (
	"errors"

	"gorm.io/gorm"
)

func (r *analyticRepository) InjectDB(db *gorm.DB) error {
	if db == nil {
		return errors.New("invalid db")
	}
	r.db = db
	return nil
}
