package main

import (
	"log"
	"strings"
)

const (
	fileExists         = "File exists"
	alreadyExists      = "already exists and is not an empty directory."
	repositoryNotFound = "Repository not found."
)

var (
	owner         = "ReSTARTR"
	keyboard      = "ergodox"
	subproject    = "ez"
	keymap        = "restartr"
	basepath      = "/Users/yoshida/src/github.com"
	toolRepo      = "https://github.com/jackhumbert/qmk_firmware"
	buildName     = keyboard + "-" + subproject + "-" + keymap
	keymapRepo    = "https://github.com/" + owner + "/qmk_firmware-" + buildName
	toolCloneTo   = basepath + "/jackhumbert/qmk_firmware"
	keymapCloneTo = basepath + "/" + owner + "/qmk_firmware-" + buildName
)

type InitCommand struct{}

func (c *InitCommand) Synopsis() string {
	return "Init keymap"
}

func (c *InitCommand) Help() string {
	return "Usage: gorgodox build [option]"
}

func (c *InitCommand) Run(args []string) int {
	var err error

	err = c.CloneTool()
	if err != nil {
		log.Fatal()
	}

	err = c.CloneKeymap()
	if err != nil {
		log.Fatal()
	}

	err = c.LinkKeymap()
	if err != nil {
		log.Fatal(err)
	}

	return 0
}

// git clone https://github.com/jackhumbert/qmk_firmware ${BASEPATH}/jackhumbert/qmk_firmware
func (c *InitCommand) CloneTool() error {
	buf, err := run("git", "clone", toolRepo, basepath+"/jackhumbert/qmk_firmware")
	if err != nil {
		if !strings.Contains(buf, alreadyExists) {
			return err
		}
	}
	return nil
}

// git clone https://github.com/${KEYMAP_REPO} ${BASEPATH}/${KEYMAP_REPO}
func (c *InitCommand) CloneKeymap() error {
	buf, err := run("git", "clone", keymapRepo, toolCloneTo)
	if err != nil {
		if !strings.Contains(buf, alreadyExists) {
			return err
		}
	}
	return nil
}

// ln -s ${BASEPATH}/${KEYMAP_REPO} \ //   ${BASEPATH}/qmk_firmware/keyboards/${KEYBOARD}/keymaps/${KEYMAP}
func (c *InitCommand) LinkKeymap() error {
	buf, err := run("ln", "-s", keymapCloneTo+"/", toolCloneTo+"/keyboards/"+keyboard+"/keymaps/"+keymap)
	if err != nil {
		if !strings.Contains(buf, fileExists) {
			return err
		}
	}
	return nil
}
