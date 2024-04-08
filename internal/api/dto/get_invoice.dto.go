package dto

type GetInvoiceResponse struct {
	Cnpj            string `json:"cnpj" example:"12345678901234"`
	Chave           string `json:"chave" example:"12345678901234567890123456789012345678901234"`
	DataEmissao     string `json:"data_emissao" example:"2022-08-01T10:00:00Z"`
	DataRecebimento string `json:"data_recebimento" example:"2022-08-01T10:00:00Z"`
} // @name GetInvoiceResponse
