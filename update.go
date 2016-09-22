package main

import "fmt"

type UpdateCommand struct{}

func (c *UpdateCommand) Synopsis() string {
	return "Update keymap"
}

func (c *UpdateCommand) Help() string {
	return "Usage: gorgodox build [option]"
}

func (c *UpdateCommand) Run(args []string) int {
	fmt.Println("build")
	return 0
}
