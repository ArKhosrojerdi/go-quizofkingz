package httpserver

import (
	"fmt"
	"quizofkingz/config"
	"quizofkingz/service/authservice"
	"quizofkingz/service/userservice"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Server struct {
	config  config.Config
	authSvc authservice.Service
	userSvc userservice.Service
}

func New(config config.Config, authSvc authservice.Service, userSvc userservice.Service) Server {
	return Server{
		config:  config,
		authSvc: authSvc,
		userSvc: userSvc,
	}
}

func (s Server) Serve() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/health-check", s.healthCheck)

	userG := e.Group("/users")
	userG.POST("/register", s.Register)
	userG.POST("/login", s.Login)
	userG.POST("/refresh", s.Refresh)
	userG.GET("/profile", s.Profile)

	err := e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port))
	if err != nil {
		return
	}
}
