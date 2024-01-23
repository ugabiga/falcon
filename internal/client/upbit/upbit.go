package upbit

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"hash"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type ErrorResp struct {
	Error struct {
		Name    float64 `json:"name"`
		Message string  `json:"message"`
	} `json:"error"`
}

type ErrorRespRaw struct {
	Error struct {
		Name    interface{} `json:"name"`
		Message string      `json:"message"`
	} `json:"error"`
}

type Client struct {
	basedURL  string
	key       string
	secret    string
	queryHash hash.Hash
}

func NewUpbitClient(key, secret string) *Client {
	return &Client{
		basedURL:  "https://api.upbit.com",
		key:       key,
		secret:    secret,
		queryHash: sha512.New(),
	}
}

func (c *Client) Accounts() ([]Account, error) {
	req, err := c.newRequest(http.MethodGet, "/v1/accounts", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	var accounts []Account
	if err = json.NewDecoder(resp.Body).Decode(&accounts); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (c *Client) TickerPublic(ctx context.Context, symbol string) (*Ticker, error) {
	params := url.Values{
		"markets": []string{symbol},
	}

	req, err := c.newPublicRequest(http.MethodGet, "/v1/ticker", params)
	if err != nil {
		log.Printf("Error creating request: %s", err.Error())
		return nil, err
	}

	var ticker []Ticker
	if err := c.do(req, &ticker); err != nil {
		return nil, err
	}

	return &ticker[0], nil
}

func (c *Client) Ticker(ctx context.Context, symbol string) (*Ticker, error) {
	params := url.Values{
		"markets": []string{symbol},
	}

	req, err := c.newRequest(http.MethodGet, "/v1/ticker", params)
	if err != nil {
		log.Printf("Error creating request: %s", err.Error())
		return nil, err
	}

	var ticker []Ticker
	if err := c.do(req, &ticker); err != nil {
		return nil, err
	}

	return &ticker[0], nil
}

func (c *Client) PlaceLongOrderAt(ctx context.Context, symbol, size, priceInKRW string) (*CreateOrderResponse, error) {
	params := url.Values{
		"market":   {symbol},
		"side":     {"bid"},
		"price":    {priceInKRW},
		"volume":   {size},
		"ord_type": {"limit"},
	}

	req, err := c.newRequest(http.MethodPost, "/v1/orders", params)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	var createOrderResponse CreateOrderResponse
	if err = json.NewDecoder(resp.Body).Decode(&createOrderResponse); err != nil {
		return nil, err
	}

	return &createOrderResponse, nil
}

func (c *Client) PlaceOrder(ctx context.Context, symbol, sizeInKRW string) (*CreateOrderResponse, error) {
	params := url.Values{
		"market":   {"KRW-" + symbol}, //Market ID
		"side":     {"bid"},           //bid(buy) or ask(sell)
		"price":    {sizeInKRW},       // Price
		"ord_type": {"price"},         // limit, price(market)
	}

	req, err := c.newRequest(http.MethodPost, "/v1/orders", params)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	var createOrderResponse CreateOrderResponse
	if err = json.NewDecoder(resp.Body).Decode(&createOrderResponse); err != nil {
		return nil, err
	}

	return &createOrderResponse, nil
}

func (c *Client) OrderChance(ctx context.Context, symbol string) (*OrderChange, error) {
	params := url.Values{
		"market": []string{symbol},
	}

	req, err := c.newRequest(http.MethodGet, "/v1/orders/chance", params)
	if err != nil {
		return nil, err
	}

	var orderChange OrderChange
	if err := c.do(req, &orderChange); err != nil {
		return nil, err
	}

	return &orderChange, nil
}

func (c *Client) OrderBook(ctx context.Context, symbol string) (*OrderBook, error) {
	params := url.Values{
		"markets": []string{symbol},
	}

	req, err := c.newRequest(http.MethodGet, "/v1/orderbook", params)
	if err != nil {
		return nil, err
	}

	var r OrderBooks
	if err := c.do(req, &r); err != nil {
		return nil, err
	}

	if len(r) < 1 {
		return nil, errors.New("no orderbook")
	}

	return &r[0], nil
}

func (c *Client) Orders(ctx context.Context, symbol string) ([]Order, error) {
	params := url.Values{
		"market": []string{symbol},
		"state":  []string{"wait"},
		"page":   []string{"1"},
		"limit":  []string{"100"},
	}

	req, err := c.newRequest(http.MethodGet, "/v1/orders", params)
	if err != nil {
		return nil, err
	}

	var orders []Order
	if err := c.do(req, &orders); err != nil {
		return nil, err
	}

	return orders, nil
}

func (c *Client) CancelOrder(ctx context.Context, uuid string) (*Order, error) {
	params := url.Values{
		"uuid": []string{uuid},
	}

	req, err := c.newRequest(http.MethodDelete, "/v1/order", params)
	if err != nil {
		return nil, err
	}

	var order Order
	if err := c.do(req, &order); err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *Client) do(req *http.Request, v interface{}) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error doing request: %s", err.Error())
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error doing request: %s", err.Error())
		return err
	}

	var result interface{}
	if err = json.Unmarshal(body, &result); err != nil {
		log.Printf("Error doing request: %s", err.Error())
		return err
	}
	log.Printf("path %v, result: %+v", req.URL.Path, result)

	switch v := result.(type) {
	case map[string]interface{}:
		// handle the case where the JSON data is an object
		if _, ok := v["error"]; ok {
			var errorResp ErrorRespRaw

			if err = json.Unmarshal(body, &errorResp); err != nil {
				log.Printf("Error doing request: %s", err.Error())
				return err
			}

			switch errorResp.Error.Name.(type) {
			case float64:
				log.Printf("Error doing request: %s", errorResp.Error.Message)
				return fmt.Errorf("error from API: %s", errorResp.Error.Message)
			case string:
				log.Printf("Error doing request: %s", errorResp.Error.Name)
				return fmt.Errorf("error from API: %s", errorResp.Error.Name)
			}
		}
	}

	if err = json.Unmarshal(body, &v); err != nil {
		log.Printf("Error doing request: %s", err.Error())
	}

	return nil
}

func (c *Client) newPublicRequest(method, url string, values url.Values) (*http.Request, error) {
	var (
		req *http.Request
		err error
	)

	// Build request
	switch method {
	case http.MethodGet, http.MethodDelete:
		req, err = http.NewRequest(method, c.basedURL+url+"?"+values.Encode(), nil)
		if err != nil {
			return nil, err
		}
	case http.MethodPost, http.MethodPut:
		req, err = http.NewRequest(method, c.basedURL+url, strings.NewReader(values.Encode()))
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("invalid method")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

func (c *Client) newRequest(method, url string, values url.Values) (*http.Request, error) {
	var (
		req *http.Request
		err error
	)

	// Build request
	switch method {
	case http.MethodGet, http.MethodDelete:
		req, err = http.NewRequest(method, c.basedURL+url+"?"+values.Encode(), nil)
		if err != nil {
			return nil, err
		}
	case http.MethodPost, http.MethodPut:
		req, err = http.NewRequest(method, c.basedURL+url, strings.NewReader(values.Encode()))
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("invalid method")
	}

	// Build token
	var token string
	token, err = c.newTokenString(values)
	if err != nil {
		return nil, err
	}

	// Set header
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+token)

	return req, nil
}

func (c *Client) newTokenString(values url.Values) (string, error) {
	claim := jwt.MapClaims{
		"access_key": c.key,
		"nonce":      uuid.New().String(),
	}

	if len(values) > 0 {
		claim["query"] = values.Encode()
		claim["query_hash"] = hex.EncodeToString(c.queryHash.Sum(nil))
		claim["query_hash_alg"] = "SHA512"

		c.queryHash.Reset()
		c.queryHash.Write([]byte(values.Encode()))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(c.secret[:]))
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(req)
}
