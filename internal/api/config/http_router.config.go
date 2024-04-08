package config

import (
	"context"
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/richhh7g/mm-api-nfe/internal/api/controller"
	"github.com/richhh7g/mm-api-nfe/internal/api/route"
	datasource "github.com/richhh7g/mm-api-nfe/internal/infra/data/invoice"
	"github.com/richhh7g/mm-api-nfe/internal/usecase/invoice"
)

type HTTPRouterConfig struct {
	ctx *context.Context
	db  *sql.DB
}

func NewHTTPRouterConfig(ctx *context.Context, db *sql.DB) ConfigBase[*chi.Mux] {
	return &HTTPRouterConfig{
		ctx: ctx,
		db:  db,
	}
}

func (c *HTTPRouterConfig) Configure() (*chi.Mux, error) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID, middleware.RealIP, middleware.Logger, middleware.Recoverer)

	invoiceDataSource := datasource.NewInvoice(c.ctx, c.db)
	route.NewInvoiceRoute(r, controller.NewInvoice(&controller.InvoiceParams{
		Ctx:                    c.ctx,
		CheckKeyExistsUseCase:  invoice.NewCheckKeyExistsUseCase(c.ctx, invoiceDataSource),
		CreateInvoiceUseCase:   invoice.NewCreateInvoiceUseCase(c.ctx, invoiceDataSource),
		GetInvoiceByKeyUseCase: invoice.NewGetInvoiceByKeyUseCase(c.ctx, invoiceDataSource),
	}))

	route.NewDocumentationRoute(c.ctx, r)
	return r, nil
}
