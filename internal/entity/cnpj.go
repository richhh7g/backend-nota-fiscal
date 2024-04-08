package entity

type CNPJ = string

func NewCNPJ(s string) CNPJ {
	return CNPJ(s)
}
