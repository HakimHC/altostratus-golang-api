package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		method := c.Request().Method
		uri := c.Request().URL.RequestURI()

		if err := next(c); err != nil {
			c.Error(err)
		}
		responseStatus := c.Response().Status
		responseStatusText := http.StatusText(responseStatus)

		log.Printf("%s %s %d %s", method, uri, responseStatus, responseStatusText)
		return nil
	}
}
