package main

import (
	"log"
	"net/http"
)

var info = map[string]Information{
	"Generate": {
		Title:       "Generate",
		Path:        "generate",
		Description: "Generates separate marksheets according tutor or class from a blank central results template.",
	},
	"Record": {
		Title:       "Record",
		Path:        "record",
		Description: "Records marks from separate marksheets submitted by tutors onto central results template.",
	},
	"Cockpit": {
		Title:       "Cockpit",
		Path:        "cockpit",
		Description: "Copies the total mark from your central results list onto Cockpit-generated CSV templates, which can then be uploaded back to Cockpit directly.",
	},
	"Analyse": {
		Title:       "Analyse",
		Path:        "analyse",
		Description: "Analyses results data. Work in progress.",
	},
}

func index(w http.ResponseWriter, r *http.Request) {
	tplDot := []Information{
		info["Generate"],
		info["Record"],
		info["Cockpit"],
		info["Analyse"],
	}
	err := tpl.ExecuteTemplate(w, "index.gohtml", tplDot)
	if err != nil {
		log.Fatal("unable to execute template - ", err)
	}
}
