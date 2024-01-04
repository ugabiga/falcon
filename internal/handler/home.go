package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/ugabiga/falcon/internal/service"
	"io"
	"log"
	"net/http"
	"time"
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
	e.GET("/", h.Index)
	e.GET("/current", h.CurrentPrice)
}

func (h HomeHandler) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "this is falcon",
	})
}

type CurrentPrice struct {
	Time struct {
		Updated    string    `json:"updated"`
		UpdatedISO time.Time `json:"updatedISO"`
		Updateduk  string    `json:"updateduk"`
	} `json:"time"`
	Disclaimer string `json:"disclaimer"`
	ChartName  string `json:"chartName"`
	Bpi        struct {
		Usd struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"USD"`
		Gbp struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"GBP"`
		Eur struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"EUR"`
	} `json:"bpi"`
}

func (h HomeHandler) CurrentPrice(c echo.Context) error {
	log.Println("getServerIP")
	resp, err := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	var respJson CurrentPrice
	err = json.Unmarshal(body, &respJson)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, respJson)
}
