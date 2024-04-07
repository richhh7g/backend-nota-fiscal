package config

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/richhh7g/mm-api-nfe/internal/api/response"
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

	r.Get("/hello-world", func(w http.ResponseWriter, _ *http.Request) {
		response.NewSuccess("Hello World! MM", http.StatusOK).Send(w)
	})

	return r, nil
}
