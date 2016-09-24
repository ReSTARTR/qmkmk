package main

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

type Option struct {
	Owner      string
	Keyboard   string
	Subproject string
	Keymap     string
	Basepath   string

	BuildName     string
	ToolRepo      string
	ToolCloneTo   string
	KeymapRepo    string
	KeymapCloneTo string
	KeymapLinkTo  string
	KeyboardDir   string
	HexDir        string
	HexFile       string
}

func (o *Option) Resolve() {
	o.BuildName = strings.Join([]string{o.Keyboard, o.Subproject, o.Keymap}, "-")
	o.ToolRepo = fmt.Sprintf("https://github.com/%s/%s", toolOwner, toolName)
	o.ToolCloneTo = filepath.Join(o.Basepath, toolOwner, toolName)
	o.KeymapRepo = "https://github.com/" + o.Owner + "/qmk_firmware-" + o.BuildName
	o.KeyboardDir = filepath.Join(o.ToolCloneTo, "keyboards", o.Keyboard)
	o.KeymapCloneTo = filepath.Join(o.Basepath, o.Owner, "qmk_firmware-"+o.BuildName)
	o.KeymapLinkTo = filepath.Join(o.ToolCloneTo, "keyboards", o.Keyboard, "keymaps", o.Keymap)
	o.HexDir = filepath.Join(o.ToolCloneTo, ".build")
	o.HexFile = strings.Replace(o.BuildName, "-", "_", -1) + ".hex"

	//fmt.Println(o.String())
}

func (o *Option) String() string {
	str, err := json.MarshalIndent(o, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	return string(str)
}
