package logger

import (
	"api-localization/logger"
	"time"

	"github.com/labstack/echo"
)

// Logger :
func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (e error) {

		start := time.Now()
		next(c)
		defer logger.Request(c.Request().Method, c.Response().Status, c.Request().RequestURI, start)

		return
	}
}
