package upbit_test

import (
	"context"
	"github.com/ugabiga/falcon/internal/app"
	"github.com/ugabiga/falcon/internal/client/upbit"
	"github.com/ugabiga/falcon/internal/common/debug"
	"log"
	"testing"
)

func TestClient_Ticker(t *testing.T) {
	tester := app.InitializeTestApplication()
	ctx := context.Background()

	client := upbit.NewUpbitClient(
		tester.Cfg.TestUpbitKey,
		tester.Cfg.TestUpbitSecret,
	)

	r, err := client.Ticker(ctx, "KRW-BTC")
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("r: %+v", debug.ToJSONStr(r))
}
