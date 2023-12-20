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
}

func NewServer(
	jwtService *service.JWTService,
	authenticationHandler *handler.AuthenticationHandler,
) *Server {
	return &Server{
		e:                     echo.New(),
		jwtService:            jwtService,
		authenticationHandler: authenticationHandler,
	}
}

func (s *Server) middleware() {
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())
	s.e.Use(s.jwtService.Middleware([]string{
		"/api/v1/auth",
	}))
}

func (s *Server) router() {
	v1 := s.e.Group("/api/v1")
	s.authenticationHandler.SetRoutes(v1)
}

func (s *Server) Run() error {
	s.middleware()
	s.router()

	s.e.Logger.Fatal(s.e.Start(":8080"))
	return nil
}
