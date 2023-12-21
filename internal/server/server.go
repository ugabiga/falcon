package server

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
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
	s.e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	//s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins:     []string{"http://localhost:3000"},
	//	AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	//	AllowCredentials: true,
	//}))
	//s.e.Use(s.jwtService.Middleware([]string{
	//	"/",
	//	"/auth/signin",
	//	"/auth/signin/:provider",
	//	"/",
	//	"/playground",
	//}))
}

func (s *Server) router() {
	s.e.Static("/static", "web/static")

	r := s.e.Group("")
	s.homeHandler.SetRoutes(r)

	r.GET("/auth/signin", s.authenticationHandler.SignInIndex)
	r.GET("/auth/signin/:provider", s.authenticationHandler.SignIn)
	r.GET("/auth/signin/:provider/callback", s.authenticationHandler.SignInCallback)
	r.GET("/auth/signout/:provider", s.authenticationHandler.SignOut)
	r.GET("/auth/protected", s.authenticationHandler.Protected,
		s.jwtService.Middleware([]string{}),
	)
}

func (s *Server) Run() error {
	s.middleware()
	s.router()

	s.e.Logger.Fatal(s.e.Start(":3000"))
	return nil
}
