package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	copyFiles = []string{
		"keymap.c",
	}

	yamlTemplate = `
build:
  steps:
    - script:
      name: build keymap
      cwd: $WERCKER_SOURCE_DIR
`
)

type Command struct {
	opt *option
}

func NewCommand(o *option) *Command {
	return &Command{o}
}

func (c Command) Help() error {
	fmt.Printf("Usage of %s\n", os.Args[0])
	flag.PrintDefaults()
	return nil
}

func (c Command) Init() error {
	return c.cloneTool()
}

func (c Command) List() error {
	return c.listKeymaps(false)
}

func (c Command) ListAvailables() error {
	return c.listAvailables()
}

func (c Command) ListHex() error {
	return c.listHex()
}

func (c Command) Create() error {
	if err := c.createKeymap(); err != nil {
		return err
	}
	return c.LinkKeymap()
}

func (c Command) Get() error {
	if c.opt.Keymap.Included {
		log.Printf("keymap is included tool in %s", filepath.Join(c.opt.Keymap.Dir, c.opt.Keymap.Name))
		return nil
	}

	if err := c.cloneKeymap(); err != nil {
		return err
	}
	return c.LinkKeymap()
}

func (c *Command) Build() error {
	return c.build()
}

func (c *Command) Install() error {
	return c.install()
}

func (c *Command) Push() error {
	// TODO: change by qmkmk.config or cli options
	return c.push("origin", "master")
}

func (c Command) listHex() error {
	fs, err := ioutil.ReadDir(c.opt.HexDir)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, fmt.Sprintf("Read dir: %s\n", c.opt.HexDir))
	for _, f := range fs {
		if !strings.HasSuffix(f.Name(), ".hex") {
			continue
		}
		fmt.Println(f.Name())
	}
	return nil
}

func (c Command) createKeymap() error {
	// create repository dir
	err := os.Mkdir(c.opt.Keymap.CloneTo, 0755)
	if err != nil && !strings.Contains(err.Error(), fileExists) {
		return err
	}
	log.Printf("[CREATE] mkdir %s\n", c.opt.Keymap.CloneTo)

	// init keymap repository
	buf, err := runIn(c.opt.Keymap.CloneTo, "git", "init")
	if err != nil {
		return err
	}
	log.Println(buf)

	// copy keymap files
	if err := c.copyWerkerYaml(); err != nil {
		return err
	}

	for _, name := range copyFiles {
		if err := c.copyKeymapFiles(name); err != nil {
			return err
		}
	}
	return nil
}

func (c Command) copyWerkerYaml() error {
	yamlPath := filepath.Join(c.opt.Keymap.CloneTo, "werker.yml")
	f, err := os.Create(yamlPath)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, yamlTemplate)
	log.Printf("[CREATE] touch %s\n", yamlPath)
	return nil
}

func (c Command) copyKeymapFiles(name string) error {
	// TODO: make source configurable
	fromPath := filepath.Join(c.opt.Keymap.Dir, "default", name)
	toPath := filepath.Join(c.opt.Keymap.CloneTo, name)

	from, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer from.Close()

	to, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	log.Printf("[CREATE] copy %s -> %s\n", fromPath, toPath)
	return err
}

func (c Command) listAvailables() error {
	// sc := NewSearchClient()
	// names, err := sc.Search(fmt.Sprintf("qmk_firmware-%s", c.opt.Keyboard))
	// if err != nil {
	// 	return err
	// }

	// for _, name := range names {
	// 	fmt.Println(name)
	// }
	return nil
}

func (c Command) cloneTool() error {
	buf, err := run("git", "clone", c.opt.ToolRepo, c.opt.ToolCloneTo)
	if err != nil {
		if !strings.Contains(buf, alreadyExists) {
			return err
		}
	}
	fmt.Printf("[INIT] cloned tool repository: %s", c.opt.ToolCloneTo)
	return nil
}

func (c Command) listKeymaps(all bool) error {
	f, err := os.Open(c.opt.Keymap.Dir)
	if err != nil {
		return err
	}

	dirs, err := f.Readdir(0)
	if err != nil {
		return err
	}
	for _, dir := range dirs {
		if dir.Mode()&os.ModeSymlink > 0 {
			fmt.Printf("%s\t %s\n", dir.Name(), c.opt.Keymap.Repo)
		} else if all {
			fmt.Printf("%-20s\t %s/tree/master/keyboards/%s/keymaps/%s/\n", dir.Name(), c.opt.ToolRepo, c.opt.Keyboard, dir.Name())
		}
	}
	return nil
}

func (c *Command) cloneKeymap() error {
	buf, err := run("git", "clone", c.opt.Keymap.Repo, c.opt.Keymap.CloneTo)
	if err != nil {
		fmt.Println(buf)
		if !strings.Contains(buf, alreadyExists) {
			return err
		}
	}
	return nil
}

func (c *Command) LinkKeymap() error {
	err := os.Symlink(c.opt.Keymap.CloneTo, c.opt.Keymap.LinkTo)
	if err != nil {
		if !strings.Contains(err.Error(), fileExists) {
			return err
		}
	}
	log.Printf("[CREATE] link %s -> %s\n", c.opt.Keymap.CloneTo, c.opt.Keymap.LinkTo)
	return nil
}

func (c *Command) build() error {
	buf, err := runIn(c.opt.ToolCloneTo, "make", c.opt.Keymap.BuildName)
	fmt.Println(buf)
	if err != nil {
		return err
	}

	hexFile := filepath.Join(c.opt.HexDir, c.opt.Keymap.HexFile)
	fmt.Printf("[BUILD] %s\n", hexFile)
	return nil
}

func (c *Command) install() error {
	buf, err := runIn(c.opt.KeyboardsDir, "make", "teensy", "KEYMAP="+c.opt.Keymap.Name)
	fmt.Println(buf)
	if err != nil {
		return err
	}
	fmt.Printf("[INSTALL] %s\n", c.opt.Keymap.BuildName)
	return nil
}

func (c *Command) push(remote, branch string) error {
	buf, err := runIn(c.opt.Keymap.Dir, "git", "push", remote, branch)
	if err != nil {
		return err
	}
	fmt.Println(buf)
	fmt.Printf("[PUSH] %s\n", c.opt.Keymap.BuildName)
	return nil
}
