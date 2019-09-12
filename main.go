package main

import (
	"github.com/brave-experiments/apollo-devops/terra"
	"github.com/mitchellh/cli"
	"io/ioutil"
	"log"
	"os"
)

var (
	Cli *cli.CLI
)

func New() *cli.CLI {
	Cli = cli.NewCLI(
		terra.StaticCliCliName,
		terra.StaticCliCliVersion,
	)
	Cli.Args = os.Args[1:]
	Cli.Commands = RegisteredCommands

	return Cli
}

func main() {
	if "" == os.Getenv("TF_LOG")  {
		log.SetOutput(ioutil.Discard)
	}

	New()

	exitStatus, err := Cli.Run()

	if nil != err {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
