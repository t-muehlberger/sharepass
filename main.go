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

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --generate server,types,spec -o pkg/api/api.gen.go --package api openapi.yml

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	dataStore, err := getStore()
	if err != nil {
		return err
	}

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

func getStore() (secrets.Store, error) {
	pgHost, usePostgres := os.LookupEnv("PG_HOST")

	if !usePostgres {
		return data.NewDiskStore("data")
	}

	pgDb, ok := os.LookupEnv("PG_DB")
	if !ok {
		pgDb = "postgres"
	}
	pgUser, ok := os.LookupEnv("PG_USER")
	if !ok {
		pgUser = "postgres"
	}
	pgPwd, ok := os.LookupEnv("PG_PWD")
	if !ok {
		pgPwd = "postgres"
	}

	return data.NewPostgresStore(pgHost, pgUser, pgPwd, pgDb)
}
