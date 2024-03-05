package repository

import (
	"context"

	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models"
	"gorm.io/gorm"
)

const (
	ClientTableName      = "clients"
	AccountTableName     = "accounts"
	TransactionTableName = "transactions"
)

type bankRepository struct {
	db *gorm.DB
}

func NewBankRepository(db *gorm.DB) bankRepository {
	return bankRepository{
		db: db,
	}
}

type BankRepository interface {
	GetStatement(ctx context.Context, clienteId int64) (models.BankStatementDomainInterface, error)
	DoTransaction(ctx context.Context, clienteId int64, bankTransaction models.BankTransactionDomainInterface) (models.BankAccountDomainInterface, error)
}
