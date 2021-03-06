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
			{Name: "add", Usage: "add message", Aliases: []string{"create", "a"}, Action: addMessage,
				Flags: append(defaultFlags, []cli.Flag{
					cli.StringFlag{Name: "text", Usage: "message text"},
					cli.StringFlag{Name: "room", Usage: "name of room to register against"}}...)},
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

func addMessage(c *cli.Context) error {
	room := c.String("room")
	if room == "" {
		return fmt.Errorf("missing room name")
	}
	text := c.String("text")
	if text == "" {
		return fmt.Errorf("missing message text")
	}
	client, err := getClient(c)
	if err != nil {
		return err
	}
	res, err := client.CreateMessage(room, text)
	if err != nil {
		return err
	}
	fmt.Printf("success id=%s\n", res.ID)
	return nil
}
