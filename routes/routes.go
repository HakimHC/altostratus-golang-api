package routes

import (
	"github.com/HakimHC/altostratus-golang-api/controllers"
	"github.com/labstack/echo/v4"
)

func AsteroidRoutes(e *echo.Echo) {
	api := e.Group("/api/v1", serverHeader)

	//api.Use(middleware.JWTMiddleware)

	api.GET("/asteroids", controllers.GetAllAsteroids)
	api.GET("/asteroids/:id", controllers.GetAsteroidById)
	api.POST("/asteroids", controllers.PostAsteroids)
	api.PATCH("/asteroids/:id", controllers.PatchAsteroid)
	api.DELETE("/asteroids/:id", controllers.DeleteAsteroidById)

	// TODO: patch method
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("X-Version", "1.0.0")
		return next(c)
	}
}
