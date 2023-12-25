package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ugabiga/falcon/internal/handler/middleware"
	"github.com/ugabiga/falcon/internal/handler/model"
)

type TaskHandler struct {
}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}

func (h *TaskHandler) SetRoutes(e *echo.Group) {
	e.GET("/tasks", h.Index)
}

type TaskIndex struct {
	Layout model.Layout
	Title  string
}

func (h *TaskHandler) Index(c echo.Context) error {
	r := RenderPage(
		c.Response().Writer,
		TaskIndex{
			Layout: middleware.ExtractLayout(c),
			Title:  "Task Page",
		},
		"/tasks/index.html",
	)

	return r
}
