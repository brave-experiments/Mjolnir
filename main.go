package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
)

var (
	app         = cli.NewApp()
	helpCommand = cli.Command{
		Action:      mainCommand, // keep track of migration progress
		Name:        "help",
		Usage:       "type help to show help",
		ArgsUsage:   " ",
		Category:    "APOLLO COMMANDS",
		Description: `type help to show help`,
	}
)

func init() {
	app.Name = os.Getenv("CLI_NAME")
	app.Version = os.Getenv("CLI_VERSION")
	app.Description = os.Getenv("CLI_DESCRIPTION")
	app.Commands = append(app.Commands, helpCommand)
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func mainCommand(ctx *cli.Context) {
	fmt.Println("Hello, there are no commands yet")
}
