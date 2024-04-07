package config

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HTTPRouterConfig struct {
	ctx *context.Context
}

func NewHTTPRouterConfig(ctx *context.Context) ConfigBase[*chi.Mux] {
	return &HTTPRouterConfig{
		ctx: ctx,
	}
}

func (c *HTTPRouterConfig) Configure() (*chi.Mux, error) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID, middleware.RealIP, middleware.Logger, middleware.Recoverer)

	return r, nil
}
