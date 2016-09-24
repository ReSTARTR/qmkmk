package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
)

const (
	binName = "qmkmk"

	toolOwner = "jackhumbert"
	toolName  = "qmk_firmware"

	fileExists         = "file exists"
	alreadyExists      = "already exists and is not an empty directory."
	repositoryNotFound = "Repository not found."
)

var (
	version string
	opt     *option

	// flag valuess
	owner      string
	basepath   string
	keyboard   string
	subproject string
	keymap     string
)

func init() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	opt = NewOption(u.HomeDir)

	flag.StringVar(&owner, "owner", "", "repository owner")
	flag.StringVar(&basepath, "basepath", "", "basepath")
	flag.StringVar(&keyboard, "keyboard", "", "keyboard name")
	flag.StringVar(&subproject, "subproject", "", "subproject name")
	flag.StringVar(&keymap, "keymap", "", "keymap name")
}

func loadFlags() {
	flag.Parse()
	if owner != "" {
		opt.Owner = owner
	}
	if basepath != "" {
		opt.Basepath = basepath
	}
	if keyboard != "" {
		opt.Keyboard = keyboard
	}
	if subproject != "" {
		opt.Subproject = subproject
	}
	if keymap != "" {
		opt.Keymap.Name = keymap
	}
}

func main() {
	loadFlags()
	opt.Resolve()
	c := NewCommand(opt)

	var err error
	var exitCode int

	args := flag.Args()
	if len(args) == 0 {
		c.Help()
		os.Exit(0)
	}

	for _, cmd := range args {
		switch cmd {
		case "init":
			err = c.Init()
		case "list":
			err = c.List()
		case "config":
			fmt.Println(opt)
		case "list-hex":
			err = c.ListHex()
		case "list-availables":
			err = c.ListAvailables()
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
			fmt.Printf("Failed to execute: %s\n", err.Error())
		}
		if exitCode != 0 {
			os.Exit(exitCode)
		}
	}
	os.Exit(exitCode)
}
