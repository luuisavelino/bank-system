package converter

import (
	"time"

	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models/repository/entity"
)

func ConvertStatementEntityToDomain(
	account entity.AccountEntity, transactions []entity.TransactionEntity,
) models.BankStatementDomainInterface {
	transactionsDomain := make([]models.Transaction, len(transactions))

	for i, transaction := range transactions {
		transactionsDomain[i].Description = transaction.Descricao.String
		transactionsDomain[i].TypeTransaction = transaction.Tipo.String
		transactionsDomain[i].Value = transaction.Valor.Int64
		transactionsDomain[i].RealizedIn = transaction.RealizadaEm.Time
	}

	accountDomain := models.Balance{
		Total:       account.Saldo.Int64,
		Limit:       account.Limite.Int64,
		ExtractDate: time.Now(),
	}

	return models.NewBankStatementDomain(accountDomain, transactionsDomain)
}

func ConvertAccountEntityToDomain(account entity.AccountEntity) models.BankAccountDomainInterface {
	return models.NewBankAccountDomain(account.Saldo.Int64, account.Limite.Int64)
}
