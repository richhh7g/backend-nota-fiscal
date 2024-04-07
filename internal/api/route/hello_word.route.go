package route

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/richhh7g/mm-api-nfe/internal/api/handler"
)

func NewHelloWorldRoutes(ctx *context.Context, router *chi.Mux) {
	router.Route("/api/v1/hello-world", func(r chi.Router) {
		r.Get("/", handler.GetV1HelloWorld)
	})
}
