package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "spark"
	app.Usage = "Spark command line utility"
	app.Version = "0.1.0"
	app.Commands = []cli.Command{
		roomCommand(),
		hookCommand(),
		messageCommand(),
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

var defaultFlags = []cli.Flag{
	cli.StringFlag{Name: "apitoken", Usage: "Spark API key", EnvVar: "SPARK_API_TOKEN"},
	cli.BoolFlag{Name: "debug", Usage: "include debug output"},
}
