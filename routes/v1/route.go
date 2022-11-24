package v1

import (
	"discusiin/controllers/users"
	mid "discusiin/middleware"
	"discusiin/routes"
	"io"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"
)

func InitRoute(payload *routes.Payload) (*echo.Echo, io.Closer) {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	mid.LogMiddleware(e)
	trace := jaegertracing.New(e, nil)

	uHandler := users.UserHandler{
		IUserServices: payload.GetUserServices(),
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"https://novalagung.com", "https://www.google.com"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT"},
		AllowedHeaders: []string{"Content-Type", "X-CSRF-Token"},
		Debug:          true,
	})
	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	api := e.Group("/api")
	v1 := api.Group("/v1")

	users := v1.Group("/users")
	users.POST("/register", uHandler.Register) // host:port/api/v1/users/signup
	users.POST("/login", uHandler.Login)       // host:port/api/v1/users/login

	return e, trace
}
