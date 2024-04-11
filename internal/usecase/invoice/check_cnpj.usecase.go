package invoice

import (
	"context"

	"github.com/richhh7g/mm-api-nfe/internal/infra/data/document"
)

var _ CheckCNPJUseCase = &CheckCNPJUseCaseImpl{}

type CheckCNPJUseCase interface {
	Exec(cnpj string) bool
}

type CheckCNPJUseCaseImpl struct {
	ctx        context.Context
	datasource document.ValidateDataSource
}

func NewCheckCNPJUseCase(ctx context.Context, datasource document.ValidateDataSource) *CheckCNPJUseCaseImpl {
	return &CheckCNPJUseCaseImpl{
		ctx:        ctx,
		datasource: datasource,
	}
}

func (u *CheckCNPJUseCaseImpl) Exec(cnpj string) bool {
	result, err := u.datasource.IsValidCNPJ(cnpj)
	if err != nil {
		return false
	}

	return result
}
