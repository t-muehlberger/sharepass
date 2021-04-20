package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/t-muehlberger/sharepass/pkg/api"
	"github.com/t-muehlberger/sharepass/pkg/assets"
	"github.com/t-muehlberger/sharepass/pkg/data"
	"github.com/t-muehlberger/sharepass/pkg/secrets"
)

//go:generate oapi-codegen --generate server,types,spec -o pkg/api/api.gen.go --package api openapi.yml

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	dataStore := data.NewMemoryStore()

	secretsSvc := secrets.Service{Store: dataStore}

	var handler api.ServerInterface = &api.Handler{Svc: secretsSvc}

	e := echo.New()
	e.Use(middleware.Logger())

	assets.RegisterSwaggerUiHandlers(e, "/swagger")
	api.RegisterSwaggerDocHandler(e, "/api/v1/swagger.json")
	api.RegisterHandlersWithBaseURL(e, handler, "/api/v1")
	assets.RegisterWebUiHandlers(e)

	e.Logger.Fatal(e.Start(":5000"))

	return nil
}
