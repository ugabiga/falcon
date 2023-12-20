package handler

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
)

type GraphHandler struct {
	resolver *handler.Server
}

func NewGraphHandler(
	resolver *handler.Server,
) *GraphHandler {
	return &GraphHandler{
		resolver: resolver,
	}
}

func (h GraphHandler) SetRoutes(e *echo.Group) {
	e.POST("/graph", h.Get)
	e.GET("/playground", h.Playground)
}

func (h GraphHandler) Get(c echo.Context) error {
	h.resolver.ServeHTTP(c.Response(), c.Request())
	return nil
}

func (h GraphHandler) Playground(c echo.Context) error {
	playground.Handler("GraphQL", "/graph").
		ServeHTTP(c.Response(), c.Request())
	return nil
}
