package main

import (
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

var tpl *template.Template

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

func init() {
	tpl = template.Must(template.ParseGlob("public/views/*gohtml"))
}

func main() {
	err := open("http://localhost:8181/")
	if err != nil {
		log.Fatal(err)
	}

	routes()
	log.Fatal(http.ListenAndServe(":8181", nil))
}
