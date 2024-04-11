package config

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/richhh7g/mm-api-nfe/internal/api/route"
	serviceRegistry "github.com/richhh7g/mm-api-nfe/internal/infra/service_registry"
)

type HTTPRouterConfig struct {
	db  *sql.DB
	ctx context.Context
}

func NewHTTPRouterConfig(ctx context.Context, db *sql.DB) ConfigBase[*chi.Mux] {
	return &HTTPRouterConfig{
		db:  db,
		ctx: ctx,
	}
}

func (c *HTTPRouterConfig) Configure() (*chi.Mux, error) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID, middleware.RealIP, middleware.Recoverer)

	serviceRegistry.NewInvoiceRoute(c.ctx, r, c.db, http.DefaultClient)
	route.NewDocumentationRoute(c.ctx, r)
	return r, nil
}
