package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/ugabiga/falcon/internal/handler/middleware"
	"github.com/ugabiga/falcon/internal/handler/model"
	"github.com/ugabiga/falcon/internal/service"
	"io"
	"net/http"
)

type HomeHandler struct {
	authenticationService *service.AuthenticationService
}

func NewHomeHandler(
	authenticationService *service.AuthenticationService,
) *HomeHandler {
	return &HomeHandler{
		authenticationService: authenticationService,
	}
}

func (h HomeHandler) SetRoutes(e *echo.Group) {
	//e.GET("/", h.Index)
	e.GET("/", h.Root)
}

type HomeIndex struct {
	Layout model.Layout
	Title  string
}

type IPResponse struct {
	IP string `json:"ip"`
}

func (h HomeHandler) Root(c echo.Context) error {
	resp, err := http.Get("http://jsonip.com")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var ipResponse IPResponse
	err = json.Unmarshal(body, &ipResponse)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ipResponse)
}

func (h HomeHandler) Index(c echo.Context) error {
	r := RenderPage(
		c.Response().Writer,
		HomeIndex{
			Layout: middleware.ExtractLayout(c),
			Title:  "Home Page",
		},
		"/index.html",
	)

	return r
}

func (h HomeHandler) Event(c echo.Context) error {
	c.Response().Header().Set("Hx-Trigger", "myEvent")

	return nil
}
