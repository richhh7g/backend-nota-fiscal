package route

import (
	"context"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func NewDocumentationRoute(ctx *context.Context, router *chi.Mux) {
	router.Get("/docs/*", httpSwagger.Handler())
}
