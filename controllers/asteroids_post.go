package controllers

import (
	"github.com/HakimHC/altostratus-golang-api/models"
	"github.com/HakimHC/altostratus-golang-api/responses"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AsteroidsPost(c echo.Context) error {
	body := new(models.AsteroidsPostDTO)

	if err := c.Bind(body); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(body); err != nil {
		return ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	listOfAsteroids, err := getAsteroidByField(map[string]string{"name": body.Name}, "Asteroids")
	if err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, err.Error())
	} else if len(*listOfAsteroids) != 0 {
		return ErrorResponse(c, http.StatusForbidden, "Duplicate asteroid name")
	}

	asteroid := models.Asteroid{
		ID:            uuid.New().String(),
		Name:          body.Name,
		Diameter:      body.Diameter,
		DiscoveryDate: body.DiscoveryDate,
		Observations:  body.Observations,
		Distances:     body.Distances,
	}

	if err := putAsteroidInDynamoDB(asteroid, "Asteroids"); err != nil {
		return ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated,
		responses.BasicResponse{
			Status:  http.StatusCreated,
			Message: "created",
			Data:    &echo.Map{"data": asteroid},
		})
}
