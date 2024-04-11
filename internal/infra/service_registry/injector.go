//go:build wireinject
// +build wireinject

package serviceRegistry

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
	"github.com/richhh7g/mm-api-nfe/internal/api/route"
	"github.com/richhh7g/mm-api-nfe/internal/infra/data/client/database/sqlc_pg"
	"github.com/richhh7g/mm-api-nfe/internal/infra/service_registry/dependency"
)

func NewInvoiceRoute(ctx context.Context, router *chi.Mux, db sqlc_pg.DBTX, httpClient *http.Client) *route.InvoiceRouteImpl {
	wire.Build(dependency.InvoiceRouteDI, dependency.DocumentValidateClientDI, dependency.InvoiceDataSourceDI, dependency.DocumentValidateDataSourceDI, dependency.CheckCNPJUseCaseDI, dependency.CheckKeyExistsUseCaseDI, dependency.CreateInvoiceUseCaseDI, dependency.GetInvoiceByKeyUseCaseDI, dependency.InvoiceControllerDI)

	return &route.InvoiceRouteImpl{}
}
