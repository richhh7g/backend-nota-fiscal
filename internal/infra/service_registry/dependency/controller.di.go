package dependency

import (
	"github.com/google/wire"
	"github.com/richhh7g/mm-api-nfe/internal/api/controller"
)

var InvoiceControllerDI = wire.NewSet(
	wire.Struct(new(controller.InvoiceParams), "*"),
	controller.NewInvoice,
	wire.Bind(new(controller.Invoice), new(*controller.InvoiceImpl)),
)
