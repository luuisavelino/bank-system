package models

type bankAccountDomain struct {
	Limit   int64
	Balance int64
}

func (sd *bankAccountDomain) GetLimit() int64 {
	return sd.Limit
}

func (sd *bankAccountDomain) GetBalance() int64 {
	return sd.Balance
}
