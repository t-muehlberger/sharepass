package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterSwaggerDocHandler(e *echo.Echo, path string) {
	e.GET(path, func(ctx echo.Context) error {
		openapi3, err := GetSwagger()
		if err != nil {
			return fmt.Errorf("failed to decode swagger doc: %w", err)
		}

		return ctx.JSON(http.StatusOK, openapi3)
	})
}
