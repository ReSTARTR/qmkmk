package main

import "fmt"

type BuildCommand struct{}

func (c *BuildCommand) Synopsis() string {
	return "Build keymap"
}

func (c *BuildCommand) Help() string {
	return "Usage: gorgodox build [option]"
}

func (c *BuildCommand) Run(args []string) int {
	fmt.Println("build")
	return 0
}
