package main

import (
	"fmt"

	"github.com/urfave/cli"

	"core/infrastructure"
)

func getClient(c *cli.Context) (*infrastructure.SparkClient, error) {
	tok := c.String("apitoken")
	if tok == "" {
		return nil, fmt.Errorf("missing apitoken")
	}
	client := infrastructure.NewSparkClient(tok)
	client.Trace = c.Bool("debug")
	return client, nil
}
