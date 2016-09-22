package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

var (
	version = "0.0.1"
)

func main() {
	c := cli.NewCLI("mkmg", version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"init": func() (cli.Command, error) {
			return &InitCommand{}, nil
		},
		"get": func() (cli.Command, error) {
			return &GetCommand{}, nil
		},
		"build": func() (cli.Command, error) {
			return &BuildCommand{}, nil
		},
		"push": func() (cli.Command, error) {
			return &PushCommand{}, nil
		},
		"install": func() (cli.Command, error) {
			return &InstallCommand{}, nil
		},
		"update": func() (cli.Command, error) {
			return &UpdateCommand{}, nil
		},
	}
	exitCode, err := c.Run()
	if err != nil {
		fmt.Println("Failed to execute: %s\n", err.Error())
	}
	os.Exit(exitCode)
}
