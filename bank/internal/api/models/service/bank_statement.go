package service

import (
	"context"

	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/pkg/logger"
	"go.uber.org/zap"
)

func (bs *bankServiceInterface) BankStatement(
	ctx context.Context, id int64,
) (models.BankStatementDomainInterface, error) {
	logger.Info("Init AddBank service",
		zap.String("journey", "AddBank"),
	)

	bankStatement, err := bs.bankRepository.GetStatement(ctx, id)

	return bankStatement, err
}
