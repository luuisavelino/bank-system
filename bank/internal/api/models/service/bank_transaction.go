package service

import (
	"context"
	"errors"

	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/pkg/logger"
	"go.uber.org/zap"
)

var cache = make(map[int64]bool)

func (bs *bankServiceInterface) BankTransaction(
	ctx context.Context, id int64, bankTransaction models.BankTransactionDomainInterface,
) (models.BankAccountDomainInterface, error) {
	logger.Info("Init BankTransaction service",
		zap.String("journey", "Transaction"),
	)

	_, ok := cache[id]
	if !ok {
		exists, err := bs.bankRepository.CheckIfClientExists(ctx, id)
		if err != nil {
			return nil, err
		}
		if !exists {
			cache[id] = false
			return nil, errors.New("client not found")
		}

		cache[id] = true
	}

	if !cache[id] {
		return nil, errors.New("client not found")
	}

	account, err := bs.bankRepository.DoTransaction(ctx, id, bankTransaction)
	return account, err
}
