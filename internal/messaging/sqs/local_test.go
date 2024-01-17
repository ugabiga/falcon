package sqs_test

import (
	"github.com/ugabiga/falcon/internal/app"
	"testing"
)

func TestLocalHandler(t *testing.T) {
	tester := app.InitializeTestApplication()
	h := tester.MessageHandler

	if err := h.Publish(); err != nil {
		t.Fatal(err)
	}

	if err := h.Subscribe(); err != nil {
		t.Fatal(err)
	}
}
