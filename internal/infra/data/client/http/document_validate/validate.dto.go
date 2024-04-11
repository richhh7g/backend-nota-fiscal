package documentValidate

type CheckDocumentBodyDTO struct {
	CNPJ string `json:"cnpj"`
}

type CheckDocumentResponseDTO struct {
	IsValid bool `json:"isValid"`
}
