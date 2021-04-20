package assets

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed swagger-ui
var SwaggerUiFiles embed.FS

func RegisterSwaggerUiHandlers(e *echo.Echo, baseURL string) {
	g := e.Group(baseURL)

	fs, err := fs.Sub(SwaggerUiFiles, "swagger-ui")
	if err != nil {
		log.Fatalf("failed to setup swagger-ui: %v", err)
	}

	g.GET("*", echo.WrapHandler(http.StripPrefix(baseURL, http.FileServer(http.FS(fs)))))
}
