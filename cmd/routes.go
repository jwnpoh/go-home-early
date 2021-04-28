package main

import (
	"net/http"
)

func routes() {
	http.Handle("/generate", http.HandlerFunc(generate))
	http.Handle("/generate/generator", http.HandlerFunc(generateCmd))
	http.Handle("/generate/generated", http.HandlerFunc(generatorPublic))

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("public/assets"))))
	http.Handle("/", http.HandlerFunc(index))
}
