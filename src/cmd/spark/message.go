package main

import (
	"fmt"

	humanize "github.com/dustin/go-humanize"
	"github.com/urfave/cli"
)

func messageCommand() cli.Command {
	return cli.Command{
		Name:    "message",
		Usage:   "interact with messages",
		Aliases: []string{"msg", "msgs"},
		Subcommands: []cli.Command{
			{Name: "list", Usage: "list messages", Aliases: []string{"ls", "l"}, Action: listMessages,
				Flags: append(defaultFlags, cli.StringFlag{Name: "room", Usage: "name of room to list messages from"})},
		}}
}

func listMessages(c *cli.Context) error {
	room := c.String("room")
	if room == "" {
		return fmt.Errorf("missing room name")
	}
	client, err := getClient(c)
	if err != nil {
		return err
	}
	messages, err := client.ListMessages(room)
	if err != nil {
		return err
	}
	for _, message := range messages {
		fmt.Printf("[%s] :%s: %s\n", message.PersonEmail, humanize.Time(message.Created), message.Text)
	}
	return nil
}
