package converter

import (
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models"
	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models/repository/entity"
)

func ConvertEntityStatementToDomain(
	balance entity.AccountEntity, transactions []entity.TransactionEntity,
) models.BankStatementDomainInterface {
	transactionsDomain := make([]models.Transaction, len(transactions))

	for i, transaction := range transactions {
		transactionsDomain[i].Description = transaction.Description.String
		transactionsDomain[i].TypeTransaction = transaction.Type.String
		transactionsDomain[i].Value = transaction.Value.Int64
		transactionsDomain[i].RealizedIn = transaction.RealizedIn.Time
	}

	balanceDomain := models.Balance{
		Total: balance.Balance.Int64,
		Limit: balance.Limit.Int64,
	}

	return models.NewBankStatementDomain(balanceDomain, transactionsDomain)
}

func ConvertEntityBalanceToDomain(balance entity.AccountEntity) models.BankAccountDomainInterface {
	return models.NewBankAccountDomain(balance.Balance.Int64, balance.Limit.Int64)
}
