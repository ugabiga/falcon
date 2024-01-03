package main

import (
	"encoding/json"
	"github.com/akrylysov/algnhsa"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
)

type IPResponse struct {
	IP string `json:"ip"`
}

func getServerIP(c echo.Context) error {
	log.Println("getServerIP")
	resp, err := http.Get("https://jsonip.io")
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	var ipResponse IPResponse
	err = json.Unmarshal(body, &ipResponse)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, ipResponse)
}

func main() {
	e := echo.New()

	e.GET("/", getServerIP)

	algnhsa.ListenAndServe(e, nil)
}
