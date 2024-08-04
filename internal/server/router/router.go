package router

import (
	"github.com/bmstu-itstech/itsreg-api/internal/server"
	"github.com/bmstu-itstech/itsreg-api/internal/server/handler"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const (
	apiPrefix = "/api/v1"
)

func Configure(
	s *server.Server,
) {
	s.Echo.Use(middleware.Logger())

	r := s.Echo.Group(apiPrefix)

	r.GET("/swagger/*", echoSwagger.WrapHandler)

	authHandler := handler.NewAuthHandler(s)
	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)

	botsHandler := handler.NewBotsHandler(s)
	r.POST("/bots/create", botsHandler.Create)
	r.POST("/bots/:id/start", botsHandler.Start)
	r.POST("/bots/:id/stop", botsHandler.Stop)
}
