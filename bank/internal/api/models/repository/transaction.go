package repository

import (
	"context"
	"fmt"

	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models/repository/entity"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models/repository/entity/converter"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/pkg/logger"
	"go.uber.org/zap"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

// DoTransaction do one transaction
func (sr bankRepository) DoTransaction(ctx context.Context, clientId int64, bankTransaction models.BankTransactionDomainInterface) (models.BankAccountDomainInterface, error) {
	logger.Info("Init GetBank repository",
		zap.String("journey", "Repository"),
	)

	account := entity.AccountEntity{}

	tx, err := sr.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(ctx)

	getAccountQuery := fmt.Sprintf(`SELECT saldo, limite FROM %s WHERE id = $1`, AccountTableName)

	err = tx.QueryRow(ctx, getAccountQuery, clientId).Scan(&account.Saldo, &account.Limite)
	if err != nil {
		return nil, err
	}

	switch bankTransaction.GetType() {
	case "d":
		account.Saldo.Int64 -= bankTransaction.GetValue()
	case "c":
		account.Saldo.Int64 += bankTransaction.GetValue()
	}

	if account.Saldo.Int64 < -account.Limite.Int64 {
		return nil, fmt.Errorf("insufficient funds")
	}

	updateTransactionQuery := fmt.Sprintf(`INSERT INTO %v (cliente_id, valor, tipo, descricao) VALUES($1, $2, $3, $4)`, TransactionTableName)
	updateAccountQuery := fmt.Sprintf(`UPDATE %v SET saldo=$1 WHERE id=$2`, AccountTableName)

	batch := &pgx.Batch{}
	batch.Queue(
		updateTransactionQuery,
		clientId,
		bankTransaction.GetValue(),
		bankTransaction.GetType(),
		bankTransaction.GetDescription(),
	)
	batch.Queue(updateAccountQuery, account.Saldo.Int64, clientId)

	br := tx.SendBatch(ctx, batch)
	_, err = br.Exec()
	if err != nil {
		return nil, err
	}

	err = br.Close()
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	logger.Info("transaction done with success",
		zap.String("journey", "Repository"),
	)

	domain := converter.ConvertAccountEntityToDomain(account)

	return domain, nil
}
