package main

import (
	"github.com/mitchellh/cli"
	"io/ioutil"
	"log"
	"os"
)

var (
	Cli *cli.CLI
)

const (
	CliName = "apollo"
)

func New() *cli.CLI {
	Cli = cli.NewCLI(
		CliName,
		os.Getenv("CLI_VERSION"),
	)
	Cli.Args = os.Args[1:]
	Cli.Commands = RegisteredCommands

	return Cli
}

func main() {
	if os.Getenv("TF_LOG") == "" {
		log.SetOutput(ioutil.Discard)
	}

	New()

	exitStatus, err := Cli.Run()

	if nil != err {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
