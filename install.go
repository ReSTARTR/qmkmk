package main

import "fmt"

type InstallCommand struct{}

func (c *InstallCommand) Synopsis() string {
	return "Install keymap"
}

func (c *InstallCommand) Help() string {
	return "Usage: gorgodox build [option]"
}

func (c *InstallCommand) Run(args []string) int {
	fmt.Println("build")
	return 0
}
