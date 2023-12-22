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
	authenticationHandler *handler.AuthenticationHandler
	homeHandler           *handler.HomeHandler
	authenticationService *service.AuthenticationService
}

func NewServer(
	authenticationHandler *handler.AuthenticationHandler,
	homeHandler *handler.HomeHandler,
	authenticationService *service.AuthenticationService,
) *Server {
	return &Server{
		e:                     echo.New(),
		authenticationHandler: authenticationHandler,
		homeHandler:           homeHandler,
		authenticationService: authenticationService,
	}
}

func (s *Server) middleware() {
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())
	s.e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	s.e.Use(s.authenticationService.JWTMiddleware(
		[]string{
			"/",
		},
		[]string{
			"/auth/signin",
		}))
	s.e.Use(falconMiddleware.LayoutMiddleware())
}

func (s *Server) router() {
	s.e.Static("/static", "web/static")

	r := s.e.Group("")
	s.homeHandler.SetRoutes(r)

	r.GET("/auth/signin", s.authenticationHandler.SignInIndex)
	r.GET("/auth/signin/:provider", s.authenticationHandler.SignIn)
	r.GET("/auth/signin/:provider/callback", s.authenticationHandler.SignInCallback)
	r.GET("/auth/signout/:provider", s.authenticationHandler.SignOut)
	r.GET("/auth/protected", s.authenticationHandler.Protected)
}

func (s *Server) Run() error {
	s.middleware()
	s.router()

	s.e.Logger.Fatal(s.e.Start(":3000"))
	return nil
}
