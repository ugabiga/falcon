package upbit

import (
	"context"
	"log"
	"testing"
)

func TestClient_Ticker(t *testing.T) {
	ctx := context.Background()
	client := NewUpbitClient("a", "a")

	ticker, err := client.Ticker(ctx, "KRW-BTC")
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("ticker: %+v", ticker)
}
