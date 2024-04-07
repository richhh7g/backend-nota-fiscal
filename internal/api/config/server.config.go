package config

import (
	"context"

	"github.com/go-chi/chi/v5"
)

type ServerConfig struct {
	ctx  *context.Context
	port string
}

func NewServerConfig(ctx *context.Context, port string) ConfigBase[*chi.Mux] {
	return &ServerConfig{
		ctx:  ctx,
		port: port,
	}
}

func (c *ServerConfig) Configure() (*chi.Mux, error) {
	httpRouterConfig := NewHTTPRouterConfig(c.ctx)
	router, err := httpRouterConfig.Configure()
	if err != nil {
		return nil, err
	}

	return router, err
}
