package service

import (
	"context"

	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models/repository"
)

func NewBankServiceInterface(bankRepository repository.BankRepository) BankServiceInterface {
	return &bankServiceInterface{
		bankRepository: bankRepository,
	}
}

type bankServiceInterface struct {
	bankRepository repository.BankRepository
}

type BankServiceInterface interface {
	BankTransaction(ctx context.Context, id int64, bankTransaction models.BankTransactionDomainInterface) error
	BankStatement(ctx context.Context, id int64) error
}
