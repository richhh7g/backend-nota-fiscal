package dto

type CreateInvoiceBody struct {
	Cnpj            string `json:"cnpj" validate:"required,len=14" example:"12345678901234" errormgs:"CNPJ inválido"`
	Chave           string `json:"chave" validate:"required,len=44" example:"12345678901234567890123456789012345678901234" errormgs:"Chave inválida"`
	DataEmissao     string `json:"data_emissao" validate:"required" example:"2022-08-01T10:00:00Z"`
	DataRecebimento string `json:"data_recebimento" validate:"required" example:"2022-08-01T10:00:00Z"`
} // @name CreateInvoiceBody

type CreateInvoiceResponse struct {
	ID              string `json:"id" format:"uuid" example:"12345678-9abc-def0-1234-56789abcdef0"`
	Cnpj            string `json:"cnpj" example:"12345678901234"`
	Chave           string `json:"chave" example:"12345678901234567890123456789012345678901234"`
	DataEmissao     string `json:"data_emissao" example:"2022-08-01T10:00:00Z"`
	DataRecebimento string `json:"data_recebimento" example:"2022-08-01T10:00:00Z"`
} // @name CreateInvoiceResponse
