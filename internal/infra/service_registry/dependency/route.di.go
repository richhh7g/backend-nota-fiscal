package dependency

import (
	"github.com/google/wire"
	"github.com/richhh7g/mm-api-nfe/internal/api/route"
)

var InvoiceRouteDI = wire.NewSet(
	route.NewInvoiceRoute,
	wire.Bind(new(route.InvoiceRoute), new(*route.InvoiceRouteImpl)),
)
