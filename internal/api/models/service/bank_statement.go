package service

import (
	"context"

	"github.com/luuisavelino/rinha-de-backend-2024-q1/pkg/logger"
	"go.uber.org/zap"
)

func (bs *bankServiceInterface) BankStatement(ctx context.Context, id int64) error {
	logger.Info("Init AddBank service",
		zap.String("journey", "AddBank"),
	)

	return nil
}
