package main

import (
	"github.com/kpango/glg"
	"os/exec"
	"runtime"
)

func open() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command(c.Browser, "--kiosk", "http://localhost:8080")
	} else {
		cmd = exec.Command(c.Browser, "--kiosk", "http://localhost:8080")
	}

	_, err := cmd.Output()
	if err != nil {
		_ = glg.Error(err)
	}
}
