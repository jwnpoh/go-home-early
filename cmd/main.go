package main

import (
	"log"
	"os/exec"
	"runtime"
)

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func main() {
	s := newServer()

	s.port = "2021"
	s.assetPath = "/css/"
	s.assetDir = "public/assets"
	s.templateDir = "public/views"

	log.Fatal(s.start())
}
