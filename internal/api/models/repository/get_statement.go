package repository

import (
	"context"

	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models/repository/entity"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models/repository/entity/converter"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/pkg/logger"
	"go.uber.org/zap"
)

const (
	BanksTableName    = "banks"
	WorkersTableName  = "workers"
	ControlsTableName = "controls"
	SchemesTableName  = "schemes"
)

// GetBank get bank by uuid
func (sr bankRepository) GetStatement(ctx context.Context, id int64) (models.BankStatementDomainInterface, error) {
	logger.Info("Init GetBank repository",
		zap.String("journey", "Repository"),
	)

	balance := entity.AccountEntity{}
	transactions := []entity.TransactionEntity{}

	tx := sr.db.Begin()

	// do the logic here

	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	domain := converter.ConvertEntityStatementToDomain(balance, transactions)

	logger.Info("Get bank with success",
		zap.String("journey", "Repository"),
	)

	return domain, nil
}
