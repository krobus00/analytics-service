package utils

import (
	"context"

	"github.com/krobus00/analytics-service/internal/constant"
	"gorm.io/gorm"
)

func NewTxContext(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, constant.KeyDBCtx, tx)
}

func GetTxFromContext(ctx context.Context, defaultTx *gorm.DB) *gorm.DB {
	txVal := ctx.Value(constant.KeyDBCtx)
	tx, ok := txVal.(*gorm.DB)
	if !ok {
		return defaultTx
	}
	return tx
}
