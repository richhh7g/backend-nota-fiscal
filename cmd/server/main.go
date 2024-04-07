package main

import (
	"context"
	"fmt"
	"net/http"

	_ "github.com/richhh7g/mm-api-nfe/docs"
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

	port, err := environment.Get("PORT")
	if err != nil {
		panic(err)
	}

	serverConfig := config.NewServerConfig(&ctx, port.(string))
	router, err := serverConfig.Configure()
	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(fmt.Sprintf(":%s", port.(string)), router)
	if err != nil {
		panic(err)
	}
}
