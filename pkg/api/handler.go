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
	s, err := h.Svc.RevealSecret(id)
	if err != nil {
		return err
	}

	sec := Secret{
		EncryptedSecret:      &s.EncryptedSecret,
		InitializationVector: &s.InitializationVector,
	}

	return ctx.JSON(http.StatusOK, sec)
}

func (h *Handler) CreateSecret(ctx echo.Context) error {
	cr := CreateSecretRequest{}
	err := ctx.Bind(&cr)
	if err != nil {
		return err
	}

	s, err := h.Svc.CreateSecret(cr.EncryptedSecret, *cr.InitializationVector, cr.TimeToLive, cr.MaxRetrievalCount)
	if err != nil {
		return err
	}

	sm := toSecretMetadata(s)

	return ctx.JSON(http.StatusCreated, sm)
}

func (h *Handler) GetSecretMetadata(ctx echo.Context, id string) error {
	s, err := h.Svc.GetSecretMetadata(id)
	if err != nil {
		return err
	}

	sm := toSecretMetadata(s)

	return ctx.JSON(http.StatusOK, sm)
}

func toSecretMetadata(s secrets.Secret) SecretMetadata {
	return SecretMetadata{
		Id:                &s.Id,
		ExpiryTime:        &s.ExpiryTime,
		MaxRetrievalCount: &s.MaxRetrievalCount,
		RetrievalCount:    &s.RetrievalCount,
	}
}
