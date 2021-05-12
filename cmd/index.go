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
    DisplayMsgInstruction: "Please select the heading for the column that you wish to generate marksheets from:",
    DisplayMsgElab: "For example, if you want to generate marksheets for each individual tutor on the team, select the teacher name (or equivalent) column. If you want to generate marksheets by class, select the class column.",
	},
	"Record": {
		Title:       "Record",
		Path:        "record",
		Description: "Records marks from separate marksheets submitted by tutors onto central results template.",
    DisplayMsgInstruction: "Please select the heading for the student name column:",
    DisplayMsgElab: "Since we are transferring the marks of each individual student from the tutor marksheets onto our central results template, we must make sure that each student's marks are recorded accurately.",
	},
	"Cockpit": {
		Title:       "Cockpit",
		Path:        "cockpit",
		Description: "Coming soon. Copies the total mark from your central results list onto Cockpit-generated CSV templates, which can then be uploaded back to Cockpit directly.",
    DisplayMsgInstruction: "Please select the heading for the student name column:",
    DisplayMsgElab: "Since we are transferring the marks of each individual student from our central results template onto the Cockpit-generated CSV files for each class, we must make sure that each student's marks are recorded accurately.",
	},
}

func index(w http.ResponseWriter, r *http.Request) {
	tplDot := []Information{
		info["Generate"],
		info["Record"],
		info["Cockpit"],
	}
	err := tpl.ExecuteTemplate(w, "index.gohtml", tplDot)
	if err != nil {
		log.Fatal("unable to execute template - ", err)
	}
}
