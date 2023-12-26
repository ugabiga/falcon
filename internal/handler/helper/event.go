package helper

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log"
)

const (
	EventShowToast = "showToast"
)

type Event map[string]interface{}

type Toast struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func NewEvent(c echo.Context, name string, data interface{}) {
	resp := Event{
		name: data,
	}
	c.Response().Header().Set("Hx-Trigger", toJSONStr(resp))
}

func NewToastEvent(c echo.Context, t, m string) {
	resp := Event{
		EventShowToast: Toast{
			Type:    t,
			Message: m,
		},
	}
	c.Response().Header().Set("Hx-Trigger", toJSONStr(resp))
}

func toJSONStr(obj interface{}) string {
	bytes, err := json.MarshalIndent(obj, "", "")
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return ""
	}

	return string(bytes)
}
