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
	homeHandler           *handler.HomeHandler
}

func NewServer(
	jwtService *service.JWTService,
	authenticationHandler *handler.AuthenticationHandler,
	homeHandler *handler.HomeHandler,
) *Server {
	return &Server{
		e:                     echo.New(),
		jwtService:            jwtService,
		authenticationHandler: authenticationHandler,
		homeHandler:           homeHandler,
	}
}

func (s *Server) middleware() {
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())
	//s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins:     []string{"http://localhost:3000"},
	//	AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	//	AllowCredentials: true,
	//}))
	//s.e.Use(s.jwtService.Middleware([]string{
	//	"/api/v1/auth",
	//	"/graph",
	//	"/playground",
	//}))
}

func (s *Server) router() {
	s.e.Static("/static", "web/static")

	r := s.e.Group("")
	s.homeHandler.SetRoutes(r)
	s.authenticationHandler.SetRoutes(r)
}

func (s *Server) Run() error {
	s.middleware()
	s.router()

	s.e.Logger.Fatal(s.e.Start(":3000"))
	return nil
}
