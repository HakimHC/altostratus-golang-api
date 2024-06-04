package controllers

import (
	"github.com/HakimHC/altostratus-golang-api/models"
	"github.com/HakimHC/altostratus-golang-api/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

func PatchAsteroid(c echo.Context) error {
	body := new(models.AsteroidsPatchDTO)
	id := c.Param("id")

	if err := c.Bind(body); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(body); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	asteroids, err := getAsteroidByField(map[string]string{"id": id}, "Asteroids")
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	asteroid := mergeAsteroid((*asteroids)[0], *body)
	err = putAsteroidInDynamoDB(asteroid, "Asteroids")
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return c.JSON(
		http.StatusOK,
		responses.BasicResponse{
			Status:  http.StatusOK,
			Message: "updated",
			Data:    &echo.Map{"data": asteroid},
		},
	)
}
