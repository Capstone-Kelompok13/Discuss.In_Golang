package v1

import (
	// "discusiin/controllers/topics"

	"discusiin/controllers/posts"
	"discusiin/controllers/topics"
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

	tHandler := topics.TopicHandler{
		ITopicServices: payload.GetTopicServices(),
	}

	pHandler := posts.PostHandler{
		IPostServices: payload.GetPostServices(),
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT"},
		AllowedHeaders: []string{"Content-Type", "X-CSRF-Token"},
		Debug:          true,
	})
	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	api := e.Group("/api")
	v1 := api.Group("/v1")

	users := v1.Group("/users")
	users.POST("/register", uHandler.Register) // host:port/api/v1/users/register
	users.POST("/login", uHandler.Login)       // host:port/api/v1/users/login

	// config := middleware.JWTConfig{
	// 	Claims:     &dto.Token{},
	// 	SigningKey: []byte(configs.Cfg.TokenSecret),
	// }

	topics := v1.Group("/topics")
	topics.GET("", tHandler.SeeAllTopics)           // host:port/api/v1/topics/
	topics.POST("/create", tHandler.CreateNewTopic) // host:port/api/v1/topics/create
	topics.GET("/:id", tHandler.SeeTopic)           // host:port/api/v1/topics/1
	topics.PUT("/:id/edit", tHandler.UpdateDescriptionTopic)
	topics.DELETE("/:id", tHandler.DeleteTopic)
	// topics.POST("/create", tHandler.CreateNewTopic, middleware.JWTWithConfig(config)) // host:port/api/v1/topics/create
	// topics.PUT("/:id/edit", tHandler.UpdateDescriptionTopic, middleware.JWTWithConfig(config))

	posts := v1.Group("/posts")
	posts.POST("/:name/create", pHandler.CreateNewPost)
	posts.GET("/:name", pHandler.SeeAllPost)
	posts.GET("/:name/:id", pHandler.SeePost)
	posts.PUT("/:name/:id/edit", pHandler.EditPost)
	posts.DELETE("/:name/:id", pHandler.DeletePost)
	return e, trace
}
