package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/t-muehlberger/sharepass/pkg/secrets"
)

var _ ServerInterface = &Handler{}

var revealSecretCounter = promauto.NewCounter(prometheus.CounterOpts{Name: "reveal_secret_ops_total"})
var createSecretCounter = promauto.NewCounter(prometheus.CounterOpts{Name: "create_secret_ops_total"})
var getMetadataCounter = promauto.NewCounter(prometheus.CounterOpts{Name: "get_metadata_ops_total"})

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

	revealSecretCounter.Inc()
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

	createSecretCounter.Inc()
	return ctx.JSON(http.StatusCreated, sm)
}

func (h *Handler) GetSecretMetadata(ctx echo.Context, id string) error {
	s, err := h.Svc.GetSecretMetadata(id)
	if err != nil {
		return err
	}

	sm := toSecretMetadata(s)

	getMetadataCounter.Inc()
	return ctx.JSON(http.StatusOK, sm)
}

func toSecretMetadata(s secrets.Secret) SecretMetadata {
	maxRetrievalCount := &s.MaxRetrievalCount
	retrievalCount := &s.RetrievalCount
	if s.AllowUnlimitedRetrieval {
		maxRetrievalCount = nil
		retrievalCount = nil
	}

	expiryTime := &s.ExpiryTime
	if s.DisableExpiryTime {
		expiryTime = nil
	}

	return SecretMetadata{
		Id:                &s.Id,
		ExpiryTime:        expiryTime,
		MaxRetrievalCount: maxRetrievalCount,
		RetrievalCount:    retrievalCount,
	}
}
