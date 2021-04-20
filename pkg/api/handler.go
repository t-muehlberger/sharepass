package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/t-muehlberger/sharepass/pkg/secrets"
)

var _ ServerInterface = &Handler{}

type Handler struct {
	Svc secrets.Service
}

func (h *Handler) RevealSecret(ctx echo.Context, id string) error {

	s := Secret{}

	return ctx.JSON(http.StatusOK, s)
}

func (h *Handler) CreateSecret(ctx echo.Context) error {
	cr := CreateSecretRequest{}
	err := ctx.Bind(&cr)
	if err != nil {
		return err
	}

	sm := SecretMetadata{}

	return ctx.JSON(http.StatusCreated, sm)
}

func (h *Handler) GetSecretMetadata(ctx echo.Context, id string) error {

	sm := SecretMetadata{}

	return ctx.JSON(http.StatusCreated, sm)
}
