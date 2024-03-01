package request

type BankTransactionRequest struct {
	Valor     int64  `json:"valor" binding:"required"`
	Tipo      string `json:"tipo" binding:"required,oneof=c d"`
	Descricao string `json:"descricao" binding:"required,min=1,max=10"`
}
