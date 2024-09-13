package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/t-muehlberger/sharepass/pkg/secrets"
)

var storedSecretsCount = promauto.NewGauge(prometheus.GaugeOpts{Name: "secrets_stored"})

func RegisterSwaggerDocHandler(e *echo.Echo, path string) {
	e.GET(path, func(ctx echo.Context) error {
		openapi3, err := GetSwagger()
		if err != nil {
			return fmt.Errorf("failed to decode swagger doc: %w", err)
		}

		return ctx.JSON(http.StatusOK, openapi3)
	})
}

func RegisterMetricsHandler(e *echo.Echo, metricsPassword string) {
	metricsUser := "metrics"
	e.GET("/metrics",
		echo.WrapHandler(promhttp.Handler()),
		middleware.BasicAuth(func(user, pass string, ctx echo.Context) (bool, error) {
			// Skip authentication if password is empty
			if metricsPassword == "" {
				return true, nil
			}
			if user == metricsUser && pass == metricsPassword {
				return true, nil
			}
			return false, nil
		}))
}

// Periodically updates the metric for the number of currently stored secrets in the background
func CollectCountBackground(store secrets.Store, ctx context.Context) {
	go func() {
		timer := time.NewTicker(1 * time.Hour)
		for ctx.Err() == nil {
			count, err := store.Count()
			if err != nil {
				fmt.Printf("failed to collect secret count: %v", err)
			}
			storedSecretsCount.Set(float64(count))
			select {
			case <-ctx.Done():
			case <-timer.C:
			}
		}
	}()
}
