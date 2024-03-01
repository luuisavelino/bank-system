package models

type bankTransactionDomain struct {
	Value       int64
	Type        string
	Description string
}

func (sd *bankTransactionDomain) GetValue() int64 {
	return sd.Value
}

func (sd *bankTransactionDomain) GetType() string {
	return sd.Type
}

func (sd *bankTransactionDomain) GetDescription() string {
	return sd.Description
}
