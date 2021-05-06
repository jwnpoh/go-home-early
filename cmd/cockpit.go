package main

import (
	"fmt"
	"log"
	"net/http"
)

func cockpit(w http.ResponseWriter, r *http.Request) {
	info := info["Cockpit"]
	err := tpl.ExecuteTemplate(w, "cockpit.gohtml", info)
	if err != nil {
		log.Fatal("unable to execute template - ", err)
	}
	fmt.Printf("Selected function: %v\n==> ", info.Title)
}
