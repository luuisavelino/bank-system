package models

import "time"

type BankStatementDomainInterface interface {
	GetTotalBalance() int64
	GetLimit() int64
	GetLastTransactions() []Transaction
	GetTransactionValue(transactionId int) int64
	GetTransactionDescription(transactionId int) string
	GetTransactionType(transactionId int) string
	GetTransactionRealizedIn(transactionId int) time.Time
}

func NewBankStatementDomain(balance Balance, transactions []Transaction) BankStatementDomainInterface {
	return &bankStatementDomain{
		Balance:      balance,
		Transactions: transactions,
	}
}

type BankTransactionDomainInterface interface {
	GetValue() int64
	GetType() string
	GetDescription() string
}

func NewBankTransactionDomain(value int64, description, typeTransaction string) BankTransactionDomainInterface {
	return &bankTransactionDomain{
		Type:        typeTransaction,
		Value:       value,
		Description: description,
	}
}

type BankAccountDomainInterface interface {
	GetLimit() int64
	GetBalance() int64
}

func NewBankAccountDomain(limit, balance int64) BankAccountDomainInterface {
	return &bankAccountDomain{
		Limit:   limit,
		Balance: balance,
	}
}
