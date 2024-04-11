package entity

type Invoice struct {
	ID         ID
	IssuedAt   string
	ReceivedAt string
	Key        Key
	CNPJ       CNPJ
}

func NewInvoice(issuedAt string, receivedAt string, key string, cnpj string) Invoice {
	return Invoice{
		ID:         NewID(),
		IssuedAt:   issuedAt,
		ReceivedAt: receivedAt,
		Key:        NewKey(key),
		CNPJ:       NewCNPJ(cnpj),
	}
}
