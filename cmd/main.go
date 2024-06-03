package main

import (
	"github.com/HakimHC/altostratus-golang-api/middleware"
	"github.com/HakimHC/altostratus-golang-api/models"
	"github.com/HakimHC/altostratus-golang-api/routes"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Validator = &models.CustomValidator{Validator: validator.New()}
	e.Use(middleware.JWTMiddleware)
	routes.AsteroidRoutes(e)
	e.Logger.Fatal(e.Start(":80"))
}
