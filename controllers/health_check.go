package controllers

import (
	"github.com/HakimHC/altostratus-golang-api/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HealthCheck(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		responses.HealthCheckResponse{
			Status: http.StatusOK,
			Reason: "Everything is healthy",
		})
}
