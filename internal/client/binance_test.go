package client_test

import (
	"context"
	"github.com/ugabiga/falcon/internal/app"
	"github.com/ugabiga/falcon/internal/client"
	"github.com/ugabiga/falcon/internal/common/debug"
	"log"
	"testing"
)

func TestBinanceClient_Balance(t *testing.T) {
	tester := app.InitializeTestApplication()
	ctx := context.Background()

	c := client.NewBinanceClient(
		tester.Cfg.TestBinanceKey,
		tester.Cfg.TestBinanceSecret,
		false,
	)

	b, err := c.Balance(ctx)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("b: %+v", debug.ToJSONStr(b))
}
