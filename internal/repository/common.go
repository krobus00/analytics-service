package repository

import (
	"fmt"

	"gorm.io/gorm"
)

func WithPagination(page int, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * limit
		return db.Limit(limit).Offset(offset)
	}
}

func WithSearch(value string, columns []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if value == "" {
			return db
		}
		query := ""
		for i, column := range columns {
			query += fmt.Sprintf("%s LIKE '%%%s%%'", column, value)
			if i < len(columns)-1 {
				query += " OR "
			}
		}
		db.Where(query)
		return db
	}
}
