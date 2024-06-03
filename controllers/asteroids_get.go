package controllers

import (
	"github.com/HakimHC/altostratus-golang-api/models"
	"github.com/HakimHC/altostratus-golang-api/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAsteroidById(c echo.Context) error {
	var id string = c.Param("id")

	asteroid, err := getAsteroidByField(map[string]string{"id": id}, "Asteroids")
	if err != nil {
		return ErrorResponse(
			c,
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	var result *models.Asteroid
	if len(*asteroid) == 0 {
		result = nil
	} else {
		result = &(*asteroid)[0]
	}

	return c.JSON(
		http.StatusOK,
		responses.BasicResponse{
			Status:  http.StatusOK,
			Message: "ok",
			Data:    &echo.Map{"data": result},
		},
	)
}

func GetAllAsteroids(c echo.Context) error {
	asteroids, err := fetchAllAsteroids("Asteroids")
	if err != nil {
		return ErrorResponse(
			c,
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return c.JSON(
		http.StatusOK,
		responses.BasicResponse{
			Status:  http.StatusOK,
			Message: "ok",
			Data:    &echo.Map{"data": asteroids},
		},
	)
}
