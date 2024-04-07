package main

import (
	"context"
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
