package repository

import (
	"context"

	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models"
	"gorm.io/gorm"
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
	GetStatement(ctx context.Context, id int64) (models.BankStatementDomainInterface, error)
}
