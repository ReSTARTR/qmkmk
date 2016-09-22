package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func run(args ...string) (string, error) {
	log.Println("[EXEC]", args) // NOTE: DEBUG

	var buf bytes.Buffer
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = &buf
	cmd.Stderr = &buf

	err := cmd.Run()
	body, _ := ioutil.ReadAll(&buf)

	return string(body), err
}
