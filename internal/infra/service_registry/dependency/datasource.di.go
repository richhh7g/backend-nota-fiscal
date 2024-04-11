package dependency

import (
	"github.com/google/wire"
	"github.com/richhh7g/mm-api-nfe/internal/infra/data/document"
	"github.com/richhh7g/mm-api-nfe/internal/infra/data/invoice"
)

var InvoiceDataSourceDI = wire.NewSet(
	invoice.NewInvoice,
	wire.Bind(new(invoice.InvoiceDataSource), new(*invoice.InvoiceDataSourceImpl)),
)

var DocumentValidateDataSourceDI = wire.NewSet(
	document.NewValidateDataSource,
	wire.Bind(new(document.ValidateDataSource), new(*document.ValidateDataSourceImpl)),
)
