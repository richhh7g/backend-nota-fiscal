package config

import (
	"context"
	"database/sql"

	"github.com/go-chi/chi/v5"
)

type ServerConfig struct {
	ctx context.Context
	db  *sql.DB
}

func NewServerConfig(ctx context.Context, db *sql.DB) ConfigBase[*chi.Mux] {
	return &ServerConfig{
		ctx: ctx,
		db:  db,
	}
}

func (c *ServerConfig) Configure() (*chi.Mux, error) {
	httpRouterConfig := NewHTTPRouterConfig(c.ctx, c.db)
	router, err := httpRouterConfig.Configure()
	if err != nil {
		return nil, err
	}

	return router, err
}
