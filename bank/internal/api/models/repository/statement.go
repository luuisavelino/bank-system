package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models/repository/entity"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models/repository/entity/converter"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/pkg/logger"
	"go.uber.org/zap"
)

// GetStatement get statement from client by id
func (sr bankRepository) GetStatement(ctx context.Context, clientId int64) (models.BankStatementDomainInterface, error) {
	logger.Info("Init GetStatement repository",
		zap.String("journey", "Repository"),
	)

	account := entity.AccountEntity{}
	transactions := []entity.TransactionEntity{}

	getAccountQuery := fmt.Sprintf(`SELECT saldo, limite FROM %s WHERE id = $1`, AccountTableName)

	err := sr.db.QueryRow(ctx, getAccountQuery, clientId).Scan(&account.Saldo, &account.Limite)
	if err != nil {
		return nil, errors.New("client not found")
	}

	getTransactionsQuery := fmt.Sprintf(`SELECT descricao, tipo, valor, realizada_em FROM %s WHERE cliente_id = $1 ORDER BY id DESC`, TransactionTableName)

	rows, err := sr.db.Query(ctx, getTransactionsQuery, clientId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t entity.TransactionEntity

		if err := rows.Scan(&t.Descricao, &t.Tipo, &t.Valor, &t.RealizadaEm); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	domain := converter.ConvertStatementEntityToDomain(account, transactions)

	logger.Info("Get bank with success",
		zap.String("journey", "Repository"),
	)

	return domain, nil
}
