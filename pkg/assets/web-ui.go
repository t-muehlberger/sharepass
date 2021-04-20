package assets

import (
	"embed"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed web-ui
var WebUiFiles embed.FS

// This handler should be registered last because it handles any route by default
func RegisterWebUiHandlers(e *echo.Echo) {
	fs, err := fs.Sub(WebUiFiles, "web-ui")
	if err != nil {
		log.Fatalf("failed to setup swagger-ui: %v", err)
	}

	// all folders containing only
	e.GET("/static/*", echo.WrapHandler(http.FileServer(http.FS(fs))))
	e.GET("/favicon.ico", echo.WrapHandler(http.FileServer(http.FS(fs))))

	// Return index.html as default for any route
	e.GET("*", func(ctx echo.Context) error {
		file, err := fs.Open("index.html")
		if err != nil {
			return err
		}
		defer file.Close()

		data, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}

		return ctx.HTMLBlob(http.StatusOK, data)
	})
}
