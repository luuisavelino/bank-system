package repository

import (
	"context"
	"errors"

	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models/repository/entity"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models/repository/entity/converter"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/pkg/logger"
	"go.uber.org/zap"
)

// GetBank get bank by uuid
func (sr bankRepository) GetStatement(ctx context.Context, clienteId int64) (models.BankStatementDomainInterface, error) {
	logger.Info("Init GetBank repository",
		zap.String("journey", "Repository"),
	)

	account := entity.AccountEntity{}
	transactions := []entity.TransactionEntity{}

	tx := sr.db.Begin()

	if err := tx.Where("id = ?", clienteId).Table(AccountTableName).First(&account).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("cliente não encontrado")
	}

	if err := tx.Table(TransactionTableName).Where("cliente_id = ?", clienteId).Find(&transactions).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	err := tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	domain := converter.ConvertStatementEntityToDomain(account, transactions)

	logger.Info("Get bank with success",
		zap.String("journey", "Repository"),
	)

	return domain, nil
}
