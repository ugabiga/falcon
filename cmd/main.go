package main

import (
	"github.com/ugabiga/falcon/internal/app"
	"github.com/ugabiga/falcon/internal/lambda"
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

var workerCmd = &cli.Command{
	Name: "worker",
	Action: func(context *cli.Context) error {
		a := app.InitializeApplication()
		return a.Worker()
	},
}

var lambdaCmd = &cli.Command{
	Name: "lambda",
	Action: func(context *cli.Context) error {
		a := app.InitializeApplication()
		return a.RunLambdaServer()
	},
}

var lambdaWorkerCmd = &cli.Command{
	Name: "lambda-worker",
	Action: func(context *cli.Context) error {
		return lambda.RunLambdaWorker()
	},
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	application := &cli.App{
		Name: "falcon",
		Commands: []*cli.Command{
			serverCmd,
			workerCmd,
			lambdaCmd,
			lambdaWorkerCmd,
		},
	}
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
