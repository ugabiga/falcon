package main

import (
	"github.com/ugabiga/falcon/internal/app"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var serverCmd = &cli.Command{
	Name: "server",
	Action: func(context *cli.Context) error {
		a := app.InitializeApplication()
		return a.RunServer()
	},
}

var lambdaServerCmd = &cli.Command{
	Name: "lambda-server",
	Action: func(context *cli.Context) error {
		a := app.InitializeApplication()
		return a.RunLambdaServer()
	},
}

var subscriberCmd = &cli.Command{
	Name: "subscriber",
	Action: func(context *cli.Context) error {
		return app.InitializeApplication().RunSubscriber()
	},
}

var publisherCmd = &cli.Command{
	Name: "publisher",
	Action: func(context *cli.Context) error {
		return app.InitializeApplication().RunPublisher()
	},
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	application := &cli.App{
		Name: "falcon",
		Commands: []*cli.Command{
			serverCmd,
			lambdaServerCmd,
			subscriberCmd,
			publisherCmd,
		},
	}
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
