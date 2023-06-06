package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LogMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "time=${time_custom}, method=${method}, uri=${uri}, status=${status}, remote_ip=${remote_ip}, bytes_in=${bytes_in}, bytes_out=${bytes_out}, latency_human=${latency_human},\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))
}
