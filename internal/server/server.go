package server

import (
	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/akrylysov/algnhsa"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ugabiga/falcon/internal/handler"
	"github.com/ugabiga/falcon/internal/handler/helper"
	falconMiddleware "github.com/ugabiga/falcon/internal/handler/middleware"
	"github.com/ugabiga/falcon/internal/service"
	"github.com/ugabiga/falcon/pkg/config"
)

type Server struct {
	e                     *echo.Echo
	cfg                   *config.Config
	authenticationService *service.AuthenticationService
	homeHandler           *handler.HomeHandler
	authenticationHandler *handler.AuthenticationHandler
	errorHandler          *handler.ErrorHandler
	graphServer           *gqlHandler.Server
}

func NewServer(
	cfg *config.Config,
	authenticationService *service.AuthenticationService,
	homeHandler *handler.HomeHandler,
	authenticationHandler *handler.AuthenticationHandler,
	errorHandler *handler.ErrorHandler,
	graphServer *gqlHandler.Server,
) *Server {
	return &Server{
		e:                     echo.New(),
		cfg:                   cfg,
		authenticationService: authenticationService,
		homeHandler:           homeHandler,
		authenticationHandler: authenticationHandler,
		errorHandler:          errorHandler,
		graphServer:           graphServer,
	}
}

func (s *Server) router() {
	s.e.HTTPErrorHandler = s.errorHandler.DebugErrorHandler

	r := s.e.Group("")
	s.homeHandler.SetRoutes(r)
	s.authenticationHandler.SetRoutes(r)

	s.e.POST("/graph", func(c echo.Context) error {
		ctx := helper.NewJWTClaimContext(c)
		r := c.Request()
		r = r.WithContext(ctx)

		s.graphServer.ServeHTTP(c.Response(), r)
		return nil
	})
}

func (s *Server) middleware() {
	s.e.Use(middleware.Recover())
	s.e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${path} (${remote_ip}) ${latency_human}\n",
		Output: s.e.Logger.Output(),
	}))
	s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{s.cfg.WebURL},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS, echo.PATCH},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	s.e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	s.e.Use(s.authenticationService.JWTMiddleware([]service.WhiteList{
		{Type: service.WhiteListTypeExact, Path: "/"},
		{Type: service.WhiteListTypeExact, Path: "/auth/signin"},
		{Type: service.WhiteListTypePrefix, Path: "/auth/signin"},
	}))
	s.e.Use(s.authenticationService.UngradedJWTMiddleware())
	s.e.Use(falconMiddleware.LayoutMiddleware())
}

func (s *Server) Run() error {
	s.middleware()
	s.router()

	s.e.Logger.Fatal(s.e.Start(":8080"))
	return nil
}
func (s *Server) RunLambda() error {
	s.middleware()
	s.router()

	algnhsa.ListenAndServe(s.e, nil)
	return nil
}
