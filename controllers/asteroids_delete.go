package controllers

import (
	"github.com/HakimHC/altostratus-golang-api/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

func DeleteAsteroidById(c echo.Context) error {
	id := c.Param("id")

	status, err := deleteAsteroid(id, "Asteroids")
	if status != http.StatusOK && err != nil {
		return ErrorResponse(
			c,
			status,
			err.Error(),
		)
	}

	return c.JSON(http.StatusNoContent, responses.BasicResponse{
		Status:  http.StatusNoContent,
		Message: "No content",
		Data:    &echo.Map{"data": nil},
	})
}
