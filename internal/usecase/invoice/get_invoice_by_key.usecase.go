package invoice

import (
	"context"

	"github.com/richhh7g/mm-api-nfe/internal/api/dto"
	"github.com/richhh7g/mm-api-nfe/internal/entity"
	"github.com/richhh7g/mm-api-nfe/internal/infra/data/invoice"
)

var _ GetInvoiceByKeyUseCase = &GetInvoiceByKeyUseCaseImpl{}

type GetInvoiceByKeyUseCase interface {
	Exec(input *string) (*dto.GetInvoiceResponse, error)
}

type GetInvoiceByKeyUseCaseImpl struct {
	ctx        *context.Context
	datasource invoice.InvoiceDataSource
}

func NewGetInvoiceByKeyUseCase(ctx *context.Context, datasource invoice.InvoiceDataSource) *GetInvoiceByKeyUseCaseImpl {
	return &GetInvoiceByKeyUseCaseImpl{
		ctx:        ctx,
		datasource: datasource,
	}
}

func (u *GetInvoiceByKeyUseCaseImpl) Exec(key *string) (*dto.GetInvoiceResponse, error) {
	invoice, err := u.datasource.FindInvoiceByKey(entity.Key(*key))
	if err != nil {
		return nil, err
	}

	return &dto.GetInvoiceResponse{
		Cnpj:            invoice.CNPJ,
		Chave:           string(invoice.Key),
		DataEmissao:     invoice.IssuedAt,
		DataRecebimento: invoice.ReceivedAt,
	}, nil
}
