package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	"core/infrastructure"
)

var defaultFlags = []cli.Flag{
	cli.StringFlag{Name: "apitoken", Usage: "Spark API key", EnvVar: "SPARK_API_TOKEN"},
	cli.BoolFlag{Name: "debug", Usage: "include debug output"},
}

func main() {
	app := cli.NewApp()
	app.Name = "spark"
	app.Usage = "Spark command line utility"
	app.Version = "0.1.0"
	app.Commands = []cli.Command{
		{
			Name: "room", Usage: "interact with rooms", Aliases: []string{"rooms", "r"}, Subcommands: []cli.Command{
				{Name: "list", Usage: "list rooms", Aliases: []string{"ls", "l"}, Action: listRooms, Flags: defaultFlags},
			},
		}, {
			Name: "hook", Usage: "interact with hooks", Aliases: []string{"hooks", "h"}, Subcommands: []cli.Command{
				{Name: "list", Usage: "list hooks", Aliases: []string{"ls", "l"}, Action: listHooks, Flags: defaultFlags},
				{Name: "add", Usage: "add hook", Aliases: []string{"create", "a"}, Action: addHook, Flags: append(defaultFlags, []cli.Flag{
					cli.StringFlag{Name: "name", Usage: "hook name"},
					cli.StringFlag{Name: "callback", Usage: "callback url"},
					cli.StringFlag{Name: "room", Usage: "name of room to register against"}}...)},
				{Name: "delete", Usage: "delete hook", Aliases: []string{"del", "rm"}, Action: deleteHook, Flags: append(defaultFlags, cli.StringFlag{Name: "name", Usage: "hook name"})},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getClient(c *cli.Context) (*infrastructure.SparkClient, error) {
	tok := c.String("apitoken")
	if tok == "" {
		return nil, fmt.Errorf("missing apitoken")
	}
	client := infrastructure.NewSparkClient(tok)
	client.Trace = c.Bool("debug")
	return client, nil
}

func listRooms(c *cli.Context) error {
	client, err := getClient(c)
	if err != nil {
		return err
	}
	rooms, err := client.ListRooms()
	if err != nil {
		return err
	}
	for _, room := range rooms {
		fmt.Printf("%s\n", room.Title)
	}
	return nil
}

func listHooks(c *cli.Context) error {
	client, err := getClient(c)
	if err != nil {
		return err
	}
	hooks, err := client.ListHooks()
	if err != nil {
		return err
	}
	for _, hook := range hooks {
		fmt.Printf("%s\n", hook.Name)
	}
	return nil
}

func addHook(c *cli.Context) error {
	name := c.String("name")
	if name == "" {
		return fmt.Errorf("missing hook name")
	}
	callbackURL := c.String("callback")
	if callbackURL == "" {
		return fmt.Errorf("missing callback URL")
	}
	room := c.String("room")

	client, err := getClient(c)
	if err != nil {
		return err
	}
	res, err := client.CreateHook(name, callbackURL, room)
	if err != nil {
		return err
	}
	fmt.Printf("success; id=%q\n", res.ID)
	return nil
}

func deleteHook(c *cli.Context) error {
	name := c.String("name")
	if name == "" {
		return fmt.Errorf("missing hook name")
	}
	client, err := getClient(c)
	if err != nil {
		return err
	}
	if err = client.DeleteHook(name); err != nil {
		return err
	}
	fmt.Println("success")
	return nil
}
