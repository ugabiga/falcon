package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ugabiga/falcon/internal/handler"
	"github.com/ugabiga/falcon/internal/service"
)

type Server struct {
	e          *echo.Echo
	jwtService *service.JWTService
}

func NewServer(
	jwtService *service.JWTService,
) *Server {
	return &Server{
		e:          echo.New(),
		jwtService: jwtService,
	}
}

func (s *Server) middleware() {
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())
	s.e.Use(s.jwtService.Middleware())
}

func (s *Server) router() {
	handler.NewAuthenticationHandler(
		s.jwtService,
	).SetRoute(s.e)
}

func (s *Server) Run() error {
	s.middleware()
	s.router()

	s.e.Logger.Fatal(s.e.Start(":8080"))
	return nil
}
