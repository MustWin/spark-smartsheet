package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func roomCommand() cli.Command {
	return cli.Command{
		Name:    "room",
		Usage:   "interact with rooms",
		Aliases: []string{"rooms", "r"},
		Subcommands: []cli.Command{
			{Name: "list", Usage: "list rooms", Aliases: []string{"ls", "l"}, Action: listRooms, Flags: defaultFlags},
		}}
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
