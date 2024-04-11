package document

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	documentValidate "github.com/richhh7g/mm-api-nfe/internal/infra/data/client/http/document_validate"
)

var _ ValidateDataSource = &ValidateDataSourceImpl{}

type ValidateDataSource interface {
	IsValidCNPJ(document string) (bool, error)
}

type ValidateDataSourceImpl struct {
	ctx    context.Context
	client documentValidate.Client
}

func NewValidateDataSource(ctx context.Context, client documentValidate.Client) *ValidateDataSourceImpl {
	return &ValidateDataSourceImpl{
		ctx:    ctx,
		client: client,
	}
}

func (d *ValidateDataSourceImpl) IsValidCNPJ(document string) (bool, error) {
	resp, err := d.client.Request(&documentValidate.ClientRequestParams{
		Url:    "/valida",
		Method: http.MethodPost,
		Data: &documentValidate.CheckDocumentBodyDTO{
			CNPJ: document,
		},
	})
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusOK {
		return false, nil
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var responseData documentValidate.CheckDocumentResponseDTO
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return false, err
	}

	return responseData.IsValid, nil
}
