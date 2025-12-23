package main

import (
	"io"
	"os/exec"
)

func ConvertToHTMLFragment(s string) string {
	var cmd *exec.Cmd = exec.Command("pandoc", "-f", "markdown", "-t", "html")
	if s == "" {
		return ""
	}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	go func() {
		defer stdin.Close()
		_, err = io.WriteString(stdin, s)
		if err != nil {
			panic(err)
		}
	}()

	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	return string(out)
}
