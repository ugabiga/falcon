package server

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ugabiga/falcon/internal/handler"
	falconMiddleware "github.com/ugabiga/falcon/internal/handler/middleware"
	"github.com/ugabiga/falcon/internal/service"
)

type Server struct {
	e                     *echo.Echo
	authenticationService *service.AuthenticationService
	homeHandler           *handler.HomeHandler
	authenticationHandler *handler.AuthenticationHandler
	userHandler           *handler.UserHandler
	errorHandler          *handler.ErrorHandler
}

func NewServer(
	authenticationService *service.AuthenticationService,
	homeHandler *handler.HomeHandler,
	authenticationHandler *handler.AuthenticationHandler,
	userHandler *handler.UserHandler,
	errorHandler *handler.ErrorHandler,
) *Server {
	return &Server{
		e:                     echo.New(),
		authenticationService: authenticationService,
		homeHandler:           homeHandler,
		authenticationHandler: authenticationHandler,
		userHandler:           userHandler,
		errorHandler:          errorHandler,
	}
}

func (s *Server) router() {
	s.e.Static("/static", "web/static")

	s.e.HTTPErrorHandler = s.errorHandler.DebugErrorHandler

	r := s.e.Group("")
	s.homeHandler.SetRoutes(r)
	s.authenticationHandler.SetRoutes(r)
	s.userHandler.SetRoutes(r)

	s.e.GET("/event", s.homeHandler.Event)
}

func (s *Server) middleware() {

	s.e.Use(middleware.Recover())
	s.e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${path} (${remote_ip}) ${latency_human}\n",
		Output: s.e.Logger.Output(),
	}))
	s.e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	s.e.Use(s.authenticationService.JWTMiddleware([]service.WhiteList{
		{Type: service.WhiteListTypeExact, Path: "/"},
		{Type: service.WhiteListTypePrefix, Path: "/auth/signin"},
		{Type: service.WhiteListTypePrefix, Path: "/static"},
	}))
	s.e.Use(s.authenticationService.UngradedJWTMiddleware())
	s.e.Use(falconMiddleware.LayoutMiddleware())
}

func (s *Server) Run() error {
	s.middleware()
	s.router()

	s.e.Logger.Fatal(s.e.Start(":3000"))
	return nil
}
