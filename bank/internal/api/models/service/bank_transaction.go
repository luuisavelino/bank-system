package service

import (
	"context"

	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/pkg/logger"
	"go.uber.org/zap"
)

func (bs *bankServiceInterface) BankTransaction(
	ctx context.Context, id int64, bankTransaction models.BankTransactionDomainInterface,
) (models.BankAccountDomainInterface, error) {
	logger.Info("Init BankTransaction service",
		zap.String("journey", "Transaction"),
	)

	account, err := bs.bankRepository.DoTransaction(ctx, id, bankTransaction)
	return account, err
}
