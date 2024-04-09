package dependency

import (
	"github.com/google/wire"
	"github.com/richhh7g/mm-api-nfe/internal/usecase/invoice"
)

var CheckKeyExistsUseCaseDI = wire.NewSet(
	invoice.NewCheckKeyExistsUseCase,
	wire.Bind(new(invoice.CheckKeyExistsUseCase), new(*invoice.CheckKeyExistsUseCaseImpl)),
)

var CreateInvoiceUseCaseDI = wire.NewSet(
	invoice.NewCreateInvoiceUseCase,
	wire.Bind(new(invoice.CreateInvoiceUseCase), new(*invoice.CreateInvoiceUseCaseImpl)),
)

var GetInvoiceByKeyUseCaseDI = wire.NewSet(
	invoice.NewGetInvoiceByKeyUseCase,
	wire.Bind(new(invoice.GetInvoiceByKeyUseCase), new(*invoice.GetInvoiceByKeyUseCaseImpl)),
)
