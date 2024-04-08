package main

import (
	_ "github.com/lib/pq"
	_ "github.com/richhh7g/mm-api-nfe/docs"

	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/richhh7g/mm-api-nfe/internal/api/config"
	"github.com/richhh7g/mm-api-nfe/internal/infra/environment"
)

func init() {
	err := environment.NewEnvLoader(nil)
	if err != nil {
		panic(err)
	}
}

// @title NFE API
// @version 1.0.0
// @description
// @contact.name Richhh7g
// @contact.url https://github.com/richhh7g
// @contact.email richhh7g@protonmail.com
// @license.name MIT
// @license.url https://github.com/richhh7g/backend-nota-fiscal/blob/main/LICENSE
// @BasePath /
func main() {
	ctx := context.Background()

	dbURL, err := environment.Get("DATABASE_URL")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", dbURL.(string))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	serverConfig := config.NewServerConfig(&ctx, db)
	router, err := serverConfig.Configure()
	if err != nil {
		panic(err)
	}

	port, err := environment.Get("PORT")
	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(fmt.Sprintf(":%s", port.(string)), router)
	if err != nil {
		panic(err)
	}
}
