package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// Dot render dot language to svg image
func Dot(src []byte) (dist []byte, err error) {
	var file *os.File
	var dotPath string

	file, err = ioutil.TempFile("", "go")
	if err != nil {
		return
	}
	defer os.Remove(file.Name())

	_, err = file.Write(src)
	if err != nil {
		return
	}

	dotPath, err = exec.LookPath("dot")
	if err != nil {
		return
	}

	cmd := exec.Command(dotPath, "-Tsvg", file.Name())

	dist, err = cmd.CombinedOutput()
	if err != nil {
		err = fmt.Errorf(string(dist))
		dist = []byte{}
	}

	return
}
