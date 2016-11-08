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
	app.Commands = []cli.Command{}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
