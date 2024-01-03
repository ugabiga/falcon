package main

import (
	"github.com/ugabiga/falcon/internal/app"
	"log"
)

func main() {
	a := app.InitializeApplication()
	if err := a.Worker(); err != nil {
		log.Fatal(err)
	}
}
