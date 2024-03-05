package response

import (
	"time"

	"github.com/luuisavelino/rinha-de-backend-2024-q1/internal/api/models"
)

type BankTransactionResponse struct {
	Limite int64 `json:"limite"`
	Saldo  int64 `json:"saldo"`
}

type Saldo struct {
	Total       int64  `json:"total"`
	DataExtrato string `json:"data_extrato"`
	Limite      int64  `json:"limite"`
}

type Transacao struct {
	Valor       int64  `json:"valor"`
	Tipo        string `json:"tipo"`
	Descricao   string `json:"descricao"`
	RealizadaEm string `json:"realizada_em"`
}

type BankStatementResponse struct {
	Saldo             Saldo       `json:"saldo"`
	UltimasTransacoes []Transacao `json:"ultimas_transacoes"`
}

func NewBankTransactionResponse(bankAccount models.BankAccountDomainInterface) BankTransactionResponse {
	return BankTransactionResponse{
		Saldo:  bankAccount.GetBalance(),
		Limite: bankAccount.GetLimit(),
	}
}

func NewBankStatementResponse(bankStatement models.BankStatementDomainInterface) BankStatementResponse {
	var ultimasTransacoes []Transacao
	for idx := range bankStatement.GetLastTransactions() {
		ultimasTransacoes = append(ultimasTransacoes, Transacao{
			Valor:       bankStatement.GetTransactionValue(idx),
			Tipo:        bankStatement.GetTransactionType(idx),
			Descricao:   bankStatement.GetTransactionDescription(idx),
			RealizadaEm: bankStatement.GetTransactionRealizedIn(idx).Format(time.RFC3339),
		})
	}

	return BankStatementResponse{
		Saldo: Saldo{
			Total:       bankStatement.GetTotalBalance(),
			Limite:      bankStatement.GetLimit(),
			DataExtrato: time.Now().Format(time.RFC3339),
		},
		UltimasTransacoes: ultimasTransacoes,
	}
}
