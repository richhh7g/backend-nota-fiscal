package invoice

import (
	"context"

	"github.com/richhh7g/mm-api-nfe/internal/entity"
	"github.com/richhh7g/mm-api-nfe/internal/infra/data/invoice"
)

var _ CheckKeyExistsUseCase = &CheckKeyExistsUseCaseImpl{}

type CheckKeyExistsUseCase interface {
	Exec(input *string) bool
}

type CheckKeyExistsUseCaseImpl struct {
	ctx        *context.Context
	datasource invoice.InvoiceDataSource
}

func NewCheckKeyExistsUseCase(ctx *context.Context, datasource invoice.InvoiceDataSource) *CheckKeyExistsUseCaseImpl {
	return &CheckKeyExistsUseCaseImpl{
		ctx:        ctx,
		datasource: datasource,
	}
}

func (u *CheckKeyExistsUseCaseImpl) Exec(key *string) bool {
	return u.datasource.FindInvoiceKeyExists(entity.Key(*key))
}
