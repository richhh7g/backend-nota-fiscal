package invoice

import (
	"context"

	"github.com/richhh7g/mm-api-nfe/internal/api/dto"
	"github.com/richhh7g/mm-api-nfe/internal/entity"
	"github.com/richhh7g/mm-api-nfe/internal/infra/data/invoice"
)

var _ CreateInvoiceUseCase = &CreateInvoiceUseCaseImpl{}

type CreateInvoiceUseCase interface {
	Exec(input *dto.CreateInvoiceBody) (*dto.CreateInvoiceResponse, error)
}

type CreateInvoiceUseCaseImpl struct {
	ctx        *context.Context
	datasource invoice.InvoiceDataSource
}

func NewCreateInvoiceUseCase(ctx *context.Context, datasource invoice.InvoiceDataSource) *CreateInvoiceUseCaseImpl {
	return &CreateInvoiceUseCaseImpl{
		ctx:        ctx,
		datasource: datasource,
	}
}

func (u *CreateInvoiceUseCaseImpl) Exec(input *dto.CreateInvoiceBody) (*dto.CreateInvoiceResponse, error) {
	invoice := entity.NewInvoice(input.DataEmissao, input.DataRecebimento, input.Chave, input.Cnpj)

	_, err := u.datasource.CreateInvoice(&invoice)
	if err != nil {
		return nil, err
	}

	return &dto.CreateInvoiceResponse{
		ID:              invoice.ID.String(),
		Cnpj:            invoice.CNPJ,
		Chave:           invoice.Key,
		DataEmissao:     invoice.IssuedAt,
		DataRecebimento: invoice.ReceivedAt,
	}, nil
}
