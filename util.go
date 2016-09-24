package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func command(args ...string) (*exec.Cmd, *bytes.Buffer) {
	log.Println("[EXEC]", args) // NOTE: DEBUG

	var buf bytes.Buffer
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = &buf // os.Stdout // &buf
	cmd.Stderr = &buf // os.Stderr // &buf

	return cmd, &buf
}

func runIn(dir string, args ...string) (string, error) {
	cmd, buf := command(args...)
	cmd.Dir = dir
	err := cmd.Run()
	body, _ := ioutil.ReadAll(buf)
	return string(body), err
}

func run(args ...string) (string, error) {
	cmd, buf := command(args...)
	err := cmd.Run()
	body, _ := ioutil.ReadAll(buf)
	return string(body), err
}
