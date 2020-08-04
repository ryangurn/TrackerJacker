package core

import (
	"bytes"
	"os/exec"
	"runtime"
)

func Command(cmd string) string {

	bash, arg := "", ""
	if runtime.GOOS == "windows" {
		bash = "powershell.exe"
		arg = "-command"
	} else if runtime.GOOS == "linux" {
		bash = "/bin/sh"
		arg = "-c"
	}

	c := exec.Command(bash, arg, cmd)
	var out bytes.Buffer
	c.Stdout = &out
	err := c.Run()
	if err != nil {
		return err.Error()
	}

	return out.String()

}