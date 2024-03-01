package response

import "time"

type BankTransactionResponse struct {
	Limite int64 `json:"limite"`
	Saldo  int64 `json:"saldo"`
}

type Saldo struct {
  Total        int64  `json:"total"`
  DataExtrato  time.Time `json:"data_extrato"`
  Limite       int64  `json:"limite"`
}

type Transacao struct {
  Valor       int64  `json:"valor"`
  Tipo        string `json:"tipo"`
  Descricao   string `json:"descricao"`
  RealizadaEm time.Time `json:"realizada_em"`
}

type BankStatementResponse struct {
  Saldo              Saldo       `json:"saldo"`
  UltimasTransacoes  []Transacao `json:"ultimas_transacoes"`
}