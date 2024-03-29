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
	v1 "github.com/ugabiga/falcon/internal/handler/v1"
	"github.com/ugabiga/falcon/internal/service"
	"github.com/ugabiga/falcon/pkg/config"
	"log"
)

type Server struct {
	e                     *echo.Echo
	cfg                   *config.Config
	authenticationService *service.AuthenticationService
	homeHandler           *handler.HomeHandler
	authenticationHandler *handler.AuthenticationHandler
	userHandler           *v1.UserHandler
	taskHandler           *v1.TaskHandler
	errorHandler          *handler.ErrorHandler
	graphServer           *gqlHandler.Server
}

func NewServer(
	cfg *config.Config,
	authenticationService *service.AuthenticationService,
	homeHandler *handler.HomeHandler,
	authenticationHandler *handler.AuthenticationHandler,
	userHandler *v1.UserHandler,
	taskHandler *v1.TaskHandler,
	errorHandler *handler.ErrorHandler,
	graphServer *gqlHandler.Server,
) *Server {
	return &Server{
		e:                     echo.New(),
		cfg:                   cfg,
		authenticationService: authenticationService,
		homeHandler:           homeHandler,
		authenticationHandler: authenticationHandler,
		userHandler:           userHandler,
		taskHandler:           taskHandler,
		errorHandler:          errorHandler,
		graphServer:           graphServer,
	}
}

func (s *Server) router() {
	s.e.HTTPErrorHandler = s.errorHandler.ErrorHandler

	r := s.e.Group("")
	s.homeHandler.SetRoutes(r)
	s.authenticationHandler.SetRoutes(r)

	apiV1Group := r.Group("/api/v1")
	s.userHandler.SetRoutes(apiV1Group)
	s.taskHandler.SetRoutes(apiV1Group)

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
	s.e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		// if request method is OPTIONS, do not log
		if c.Request().Method == "OPTIONS" {
			return
		}

		if c.Request().Method != "GET" {
			log.Printf("request:" + c.Path() + ":" + string(reqBody))
		}

		log.Printf("response:" + c.Path() + ":" + string(resBody))
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
		{Type: service.WhiteListTypeExact, Path: "/current"},
		{Type: service.WhiteListTypeExact, Path: "/auth/signin"},
		{Type: service.WhiteListTypePrefix, Path: "/auth/signin"},
	}))
	s.e.Use(s.authenticationService.UngradedJWTMiddleware())
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
