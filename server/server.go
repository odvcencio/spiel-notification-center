package server

import (
	"github.com/labstack/echo"
)

// CreateAndListen creates main echo server and
// starts listening on specific port
func CreateAndListen() {
	e := echo.New()

	// Connect logger
	e.Use(connectLogger())

	// Registering routes
	registerRoutes(e)

	e.Start(":8080")
}
