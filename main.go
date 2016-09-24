package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

const (
	binName   = "qmkmk"
	toolOwner = "jackhumbert"
	toolName  = "qmk_firmware"

	fileExists         = "File exists"
	alreadyExists      = "already exists and is not an empty directory."
	repositoryNotFound = "Repository not found."
)

var (
	version string
	option  *Option
)

func init() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	option = &Option{}

	flag.StringVar(&option.Owner, "owner", "ReSTARTR", "repository owner")
	flag.StringVar(&option.Keyboard, "keyboard", "ergodox", "keyboard name")
	flag.StringVar(&option.Subproject, "subproject", "ez", "subproject name")
	flag.StringVar(&option.Keymap, "keymap", "restartr", "keymap name")
	flag.StringVar(&option.Basepath, "basepath", filepath.Join(u.HomeDir, "src", "github.com"), "basepath")
}

func main() {
	flag.Parse()
	option.Resolve()

	c := NewCommand(option)

	var err error
	var exitCode int

	if len(os.Args) == 1 {
		c.Help()
		os.Exit(0)
	}

	for _, cmd := range os.Args[1:] {
		switch cmd {
		case "init":
			err = c.Init()
		case "get":
			err = c.Get()
		case "build":
			err = c.Build()
		case "install":
			err = c.Install()
		case "version":
			fmt.Println(version)
		default:
			fmt.Printf("[ERROR] Undefined command: %s\n", cmd)
			c.Help()
			os.Exit(1)
		}
		if err != nil {
			fmt.Println("Failed to execute: %s\n", err.Error())
		}
		if exitCode != 0 {
			os.Exit(exitCode)
		}
	}
	os.Exit(exitCode)
}
