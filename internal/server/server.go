package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ugabiga/falcon/internal/handler"
	"github.com/ugabiga/falcon/internal/service"
)

type Server struct {
	e                     *echo.Echo
	jwtService            *service.JWTService
	authenticationHandler *handler.AuthenticationHandler
	graphHandler          *handler.GraphHandler
}

func NewServer(
	jwtService *service.JWTService,
	authenticationHandler *handler.AuthenticationHandler,
	graphHandler *handler.GraphHandler,
) *Server {
	return &Server{
		e:                     echo.New(),
		jwtService:            jwtService,
		authenticationHandler: authenticationHandler,
		graphHandler:          graphHandler,
	}
}

func (s *Server) middleware() {
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())
	s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))
	s.e.Use(s.jwtService.Middleware([]string{
		"/api/v1/auth",
		"/graph",
		"/playground",
	}))
}

func (s *Server) router() {
	root := s.e.Group("")
	s.graphHandler.SetRoutes(root)

	v1 := s.e.Group("/api/v1")
	s.authenticationHandler.SetRoutes(v1)
}

func (s *Server) Run() error {
	s.middleware()
	s.router()

	s.e.Logger.Fatal(s.e.Start(":8080"))
	return nil
}
