package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Keymap struct {
	Name      string `json:"name"`
	BuildName string `json:"build_name"` // <keyboard>-<subproject>-<keymap>
	Dir       string `json:"dir"`
	Repo      string `json:"repo"`
	CloneTo   string `json:"clone_to"`
	LinkTo    string `json:"link_to"`
	HexFile   string `json:"hex_file"`
	Included  bool   `json:"-"`
}

type option struct {
	HomeDir    string `json:"home_dir"`
	Owner      string `json:"owner"`
	Keyboard   string `json:"keyboard"`
	Subproject string `json:"subproject"`
	Basepath   string `json:"basepath"`

	ToolRepo     string  `json:"tool_repo"`
	ToolCloneTo  string  `json:"tool_clone_to"`
	KeyboardsDir string  `json:"keyboards_dir"`
	HexDir       string  `json:"hex_dir"`
	Keymap       *Keymap `json:"keymap"`
}

func NewOption(homeDir string) *option {
	opt, err := LoadConfig(homeDir)
	if err != nil {
		log.Fatal(err)
	}
	return opt
}

func LoadConfig(homeDir string) (*option, error) {
	opt := option{
		HomeDir: homeDir,
		Keymap:  &Keymap{},
	}

	path := filepath.Join(homeDir, ".config", binName+".json")

	cf, err := os.Open(path)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return &opt, nil
		}
		return nil, err
	}

	buf, err := ioutil.ReadAll(cf)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(buf, &opt); err != nil {
		return nil, err
	}

	if opt.HomeDir == "" {
		opt.HomeDir = homeDir
	}
	return &opt, nil
}

func (o *option) Resolve() {
	if o.Basepath == "" {
		o.Basepath = filepath.Join(o.HomeDir, "src", "github.com")
	}

	if o.Keymap.Name == "" {
		o.Keymap.Name = "default"
	}

	o.ToolRepo = fmt.Sprintf("https://github.com/%s/%s", toolOwner, toolName)
	o.ToolCloneTo = filepath.Join(o.Basepath, toolOwner, toolName)
	o.HexDir = filepath.Join(o.ToolCloneTo, ".build")
	o.KeyboardsDir = filepath.Join(o.ToolCloneTo, "keyboards", o.Keyboard)

	o.Keymap.BuildName = strings.Join([]string{o.Keyboard, o.Subproject, o.Keymap.Name}, "-")
	o.Keymap.Repo = "https://github.com/" + o.Owner + "/qmk_firmware-" + o.Keymap.BuildName
	o.Keymap.Dir = filepath.Join(o.ToolCloneTo, "keyboards", o.Keyboard, "keymaps")
	o.Keymap.HexFile = strings.Replace(o.Keymap.BuildName, "-", "_", -1) + ".hex"

	if _, err := os.Open(o.ToolCloneTo); err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			// tools has not been initialized
			return
		}
		log.Fatal(err.Error())
	}

	dirs, err := ioutil.ReadDir(o.Keymap.Dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, dir := range dirs {
		if dir.Name() == o.Keymap.Name && dir.IsDir() {
			o.Keymap.Included = true
			break
		}
	}
	if !o.Keymap.Included {
		o.Keymap.CloneTo = filepath.Join(o.Basepath, o.Owner, "qmk_firmware-"+o.Keymap.BuildName)
		o.Keymap.LinkTo = filepath.Join(o.Keymap.Dir, o.Keymap.Name)
	}
}

func (o *option) String() string {
	str, err := json.MarshalIndent(o, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	return string(str)
}
