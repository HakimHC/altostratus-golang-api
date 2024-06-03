package middleware

import (
	"fmt"
	"github.com/HakimHC/altostratus-golang-api/config"
	"github.com/HakimHC/altostratus-golang-api/controllers"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return controllers.ErrorResponse(
				c,
				http.StatusUnauthorized,
				"missing authorization header",
			)
		}

		splitToken := strings.Split(authHeader, " ")

		if len(splitToken) != 2 || splitToken[0] != "Bearer" {
			return controllers.ErrorResponse(
				c,
				http.StatusUnauthorized,
				"invalid authorization header format",
			)
		}

		jwtToken := splitToken[1]
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(config.JwtSecret), nil
		})

		if err != nil || !token.Valid {
			return controllers.ErrorResponse(
				c,
				http.StatusUnauthorized,
				"invalid token",
			)
		}
		return next(c)
	}
}
