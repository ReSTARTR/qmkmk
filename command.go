package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Command struct {
	opt *Option
}

func NewCommand(o *Option) *Command {
	return &Command{o}
}

func (c Command) Help() error {
	fmt.Printf("Usage of %s\n", os.Args[0])
	flag.PrintDefaults()
	return nil
}

func (c Command) Init() error {
	return c.CloneTool()
}

func (c Command) Get() error {
	if err := c.CloneKeymap(); err != nil {
		return err
	}
	return c.LinkKeymap()
}

// git clone https://github.com/jackhumbert/qmk_firmware ${BASEPATH}/jackhumbert/qmk_firmware
func (c Command) CloneTool() error {
	buf, err := run("git", "clone", c.opt.ToolRepo, c.opt.ToolCloneTo)
	if err != nil {
		if !strings.Contains(buf, alreadyExists) {
			return err
		}
	}
	return nil
}

// git clone https://github.com/${KEYMAP_REPO} ${BASEPATH}/${KEYMAP_REPO}
func (c *Command) CloneKeymap() error {
	buf, err := run("git", "clone", c.opt.KeymapRepo, c.opt.KeymapCloneTo)
	if err != nil {
		fmt.Println(buf)
		if !strings.Contains(buf, alreadyExists) {
			return err
		}
	}
	return nil
}

// ln -s ${BASEPATH}/${KEYMAP_REPO} \ //   ${BASEPATH}/qmk_firmware/keyboards/${KEYBOARD}/keymaps/${KEYMAP}
func (c *Command) LinkKeymap() error {
	buf, err := run("ln", "-s", c.opt.KeymapCloneTo+"/", c.opt.KeymapLinkTo)
	if err != nil {
		if !strings.Contains(buf, fileExists) {
			return err
		}
	}
	return nil
}

func (c *Command) Build() error {
	buf, err := runIn(c.opt.ToolCloneTo, "make", c.opt.BuildName)
	fmt.Println(buf)
	if err != nil {
		return err
	}

	hexFile := filepath.Join(c.opt.HexDir, c.opt.HexFile)
	fmt.Printf("[BUILD] %s\n", hexFile)
	return nil
}

func (c *Command) Install() error {
	buf, err := runIn(c.opt.KeyboardDir, "make", "teensy", "KEYMAP="+c.opt.Keymap)
	fmt.Println(buf)
	if err != nil {
		return err
	}
	fmt.Printf("[INSTALL] %s\n", c.opt.BuildName)
	return nil
}
