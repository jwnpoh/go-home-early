package main

import (
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatal("unable to execute template - ", err)
	}
}

func css(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/css/", http.FileServer(http.Dir("public/assets")))
}

func routes() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("public/assets"))))
	// http.Handle("/css/", http.HandlerFunc(css))
	http.Handle("/generateCmd", http.HandlerFunc(generateCmd))
	// http.Handle("/upload", http.HandlerFunc(uploadSingle))
	// http.Handle("/uploadMultiple", http.HandlerFunc(uploadMultiple))
	http.Handle("/generate", http.HandlerFunc(generate))

	http.Handle("/", http.HandlerFunc(index))
}
