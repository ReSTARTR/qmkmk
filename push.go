package main

import "fmt"

type PushCommand struct{}

func (c *PushCommand) Synopsis() string {
	return "Push keymap"
}

func (c *PushCommand) Help() string {
	return "Usage: gorgodox build [option]"
}

func (c *PushCommand) Run(args []string) int {
	fmt.Println("build")
	return 0
}
