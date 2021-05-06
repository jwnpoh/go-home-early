package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const startMsg = `
Go Home Early
Author: Joel Poh
 2021 National Junior College

==> Started server, listening on port %v....
==> `

var tpl *template.Template

type server struct {
	port        string
	assetPath   string
	assetDir    string
	templateDir string
}

func (s *server) start() error {
	err := open("http://localhost:" + s.port)
	if err != nil {
		return err
	}

	fmt.Printf(startMsg, s.port)

	s.parseTemplates()
	s.serveStatic()
	s.router()
	err = http.ListenAndServe(":"+s.port, nil)
	if err != nil {
		return err
	}
	return nil
}

func newServer() *server {
	var server server
	return &server
}

func (s *server) serveStatic() {
	http.Handle(s.assetPath, http.StripPrefix(s.assetPath, http.FileServer(http.Dir(s.assetDir))))
}

func (s *server) router() {
	http.HandleFunc("/", index)
	http.HandleFunc("/generate", generate)
	http.HandleFunc("/generate/upload", generateUpload)
	http.HandleFunc("/serveFile", serveFile)
	http.HandleFunc("/record", record)
	http.HandleFunc("/record/upload", recordUpload)
	http.HandleFunc("/cockpit", cockpit)
}

func (s *server) parseTemplates() {
	templates := filepath.Join(s.templateDir, "*gohtml")
	tpl = template.Must(template.ParseGlob(templates))
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	defer os.Exit(0)
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	filename := r.Form.Get("download")
	rmdir := r.Form.Get("remove")

	defer os.RemoveAll(rmdir)

	fmt.Printf("File ready for download. Cleaning up temporary files....\n==> ")
	fmt.Println("Done!")

	filenamebase := filepath.Base(filename)
	w.Header().Set("Content-Disposition", "attachment; filename="+filenamebase)
	http.ServeFile(w, r, filename)
}
