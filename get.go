package main

import "fmt"

type GetCommand struct{}

func (c *GetCommand) Synopsis() string {
	return "Get keymap"
}

func (c *GetCommand) Help() string {
	return "Usage: gorgodox build [option]"
}

func (c *GetCommand) Run(args []string) int {
	fmt.Println("build")
	return 0
}
