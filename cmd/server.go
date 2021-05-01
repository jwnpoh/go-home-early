package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

type server struct {
	port        string
	assetPath   string
	assetDir    string
	templateDir string
}

func (s *server) start() error {
	err := open("http://localhost" + s.port)
	if err != nil {
		return err
	}

	fmt.Println("Started server. Listening at port 2021")

	s.parseTemplates()
	s.serveStatic()
	s.router()
	err = http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *server) serveStatic() {
	http.Handle(s.assetPath, http.StripPrefix(s.assetPath, http.FileServer(http.Dir(s.assetDir))))
}

func (s *server) router() {
	http.HandleFunc("/", index)
	http.HandleFunc("/generate", generate)
	http.HandleFunc("/generate/generator", generateCmd)
	http.HandleFunc("/generate/generated", generatorPublic)
	http.HandleFunc("/serveFile", serveFile)
}

func (s *server) parseTemplates() {
	templates := filepath.Join(s.templateDir, "*gohtml")
	tpl = template.Must(template.ParseGlob(templates))
}

func newServer() *server {
	var server server
	return &server
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	filename := r.Form.Get("download")
	rm := r.Form.Get("remove")

	fmt.Println(filename)

	defer os.RemoveAll(rm)
	defer os.Remove(filename)

	filenamebase := filepath.Base(filename)
	w.Header().Set("Content-Disposition", "attachment; filename="+filenamebase)
	http.ServeFile(w, r, filename)
}
