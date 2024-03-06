package models

import "time"

type Balance struct {
	Total         int64
	Limit         int64
	ExtractDate   time.Time
}

type Transaction struct {
	Value           int64
	Description     string
	TypeTransaction string
	RealizedIn      time.Time
}

type bankStatementDomain struct {
	Balance      Balance
	Transactions []Transaction
}

func (sd *bankStatementDomain) GetTotalBalance() int64 {
	return sd.Balance.Total
}

func (sd *bankStatementDomain) GetLimit() int64 {
	return sd.Balance.Limit
}

func (sd *bankStatementDomain) GetExtractDate() time.Time {
	return sd.Balance.ExtractDate
}

func (sd *bankStatementDomain) GetLastTransactions() []Transaction {
	return sd.Transactions
}

func (sd *bankStatementDomain) GetTransactionValue(transactionId int) int64 {
	return sd.Transactions[transactionId].Value
}

func (sd *bankStatementDomain) GetTransactionDescription(transactionId int) string {
	return sd.Transactions[transactionId].Description
}

func (sd *bankStatementDomain) GetTransactionType(transactionId int) string {
	return sd.Transactions[transactionId].TypeTransaction
}

func (sd *bankStatementDomain) GetTransactionRealizedIn(transactionId int) time.Time {
	return sd.Transactions[transactionId].RealizedIn
}
