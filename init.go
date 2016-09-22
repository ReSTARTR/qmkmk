package main

import "fmt"

type InitCommand struct{}

func (c *InitCommand) Synopsis() string {
	return "Init keymap"
}

func (c *InitCommand) Help() string {
	return "Usage: gorgodox build [option]"
}

func (c *InitCommand) Run(args []string) int {
	fmt.Println("build")
	return 0
}
