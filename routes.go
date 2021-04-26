package main

import (
	"net/http"
)

func routes() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("public/assets"))))
	http.Handle("/generator", http.HandlerFunc(generateCmd))
	http.Handle("/generate", http.HandlerFunc(generate))

	http.Handle("/", http.HandlerFunc(index))
}
