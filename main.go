package main

import (
	"github.com/BSick7/envoy/command"
	"github.com/mitchellh/cli"
	"log"
	"os"
)

var Version string

func main() {
	c := cli.NewCLI("envoy", Version)
	c.Args = os.Args[1:]

	metaPtr := &command.Meta{
		Ui: &cli.BasicUi{
			Writer:      os.Stdout,
			ErrorWriter: os.Stderr,
		},
		Version: Version,
	}
	meta := *metaPtr

	c.Commands = map[string]cli.CommandFactory{
		"backup": func() (cli.Command, error) {
			return &command.BackupCommand{
				Meta: meta,
			}, nil
		},
		"restore": func() (cli.Command, error) {
			return &command.RestoreCommand{
				Meta: meta,
			}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
