package main

import (
	"log"
	"net/http"
)

type Commands struct {
	Title       string
	Path        string
	Description string
}

var cmds = []Commands{
	{
		Title:       "Generate",
		Path:        "generate",
		Description: "Generates separate marksheets according tutor or class from a blank central results template.",
	},
	{
		Title:       "Record",
		Path:        "record",
		Description: "Records marks from separate marksheets submitted by tutors onto central results template.",
	},
	{
		Title:       "Cockpit",
		Path:        "cockpit",
		Description: "Copies the total mark from your central results list onto Cockpit-generated CSV templates, which can then be uploaded back to Cockpit directly.",
	},
	{
		Title:       "Analyse",
		Path:        "analyse",
		Description: "Analyses results data. Work in progress.",
	},
}

func index(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "index.gohtml", cmds)
	if err != nil {
		log.Fatal("unable to execute template - ", err)
	}
}
