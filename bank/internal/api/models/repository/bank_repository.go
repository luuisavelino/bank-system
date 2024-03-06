package repository

import (
	"context"

	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

const (
	ClientTableName      = "clients"
	AccountTableName     = "accounts"
	TransactionTableName = "transactions"
)

type bankRepository struct {
	db *pgxpool.Pool
}

func NewBankRepository(db *pgxpool.Pool) bankRepository {
	return bankRepository{
		db: db,
	}
}

type BankRepository interface {
	GetStatement(ctx context.Context, clientId int64) (models.BankStatementDomainInterface, error)
	DoTransaction(ctx context.Context, clientId int64, bankTransaction models.BankTransactionDomainInterface) (models.BankAccountDomainInterface, error)
	CheckIfClientExists(ctx context.Context, clientId int64) (bool, error)
}
