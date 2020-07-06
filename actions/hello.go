package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HelloHandler
func HelloHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("hello.html"))
}
