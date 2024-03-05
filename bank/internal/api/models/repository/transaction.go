package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models/repository/entity"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models/repository/entity/converter"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/pkg/logger"
	"go.uber.org/zap"
)

// DoTransaction do one transaction
func (sr bankRepository) DoTransaction(ctx context.Context, clienteId int64, bankTransaction models.BankTransactionDomainInterface) (models.BankAccountDomainInterface, error) {
	logger.Info("Init GetBank repository",
		zap.String("journey", "Repository"),
	)

	account := entity.AccountEntity{}

	tx := sr.db.Begin()

	if err := tx.Where("id = ?", clienteId).Table(AccountTableName).First(&account).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("cliente não encontrado")
	}

	creditLimit := account.Limite.Int64
	account.Saldo.Int64 -= bankTransaction.GetValue()

	if account.Saldo.Int64 < -creditLimit {
		tx.Rollback()
		return nil, errors.New("saldo insuficiente")
	}

	if err := tx.Table(AccountTableName).Save(&account).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	transaction := entity.TransactionEntity{
		ClienteID:   sql.NullInt64{Int64: clienteId, Valid: true},
		Valor:       sql.NullInt64{Int64: -bankTransaction.GetValue(), Valid: true},
		Tipo:        sql.NullString{String: bankTransaction.GetType(), Valid: true},
		Descricao:   sql.NullString{String: bankTransaction.GetDescription(), Valid: true},
		RealizadaEm: sql.NullTime{Time: time.Now(), Valid: true},
	}

	if err := tx.Table(TransactionTableName).Create(&transaction).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	logger.Info("transaction done with success",
		zap.String("journey", "Repository"),
	)

	account.Saldo = sql.NullInt64{Int64: account.Saldo.Int64, Valid: true}

	domain := converter.ConvertAccountEntityToDomain(account)

	return domain, nil
}
