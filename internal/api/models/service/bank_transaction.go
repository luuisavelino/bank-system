package service

import (
	"context"

	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/pkg/logger"
	"go.uber.org/zap"
)

func (bs *bankServiceInterface) BankTransaction(
	ctx context.Context, id int64, bankTransaction models.BankTransactionDomainInterface,
) error {
	logger.Info("Init RemoveBank service",
		zap.String("journey", "RemoveBank"),
	)

	return nil
}
