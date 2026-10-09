package httpserver

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type HealthCheckResponse struct {
	Message string `json:"message"`
}

func (s Server) healthCheck(c *echo.Context) error {
	m := &HealthCheckResponse{Message: "ok"}
	return c.JSON(http.StatusOK, m)
}
