package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

var t csvData

func generate(w http.ResponseWriter, r *http.Request) {
	info := info["Generate"]
	err := tpl.ExecuteTemplate(w, "generate.gohtml", info)
	if err != nil {
		log.Fatal("unable to execute template - ", err)
	}
	fmt.Printf("Selected function: %v\n==> ", info.Title)
}

func generateUpload(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		tmpFile := uploadSingle(w, r)
		// Take file bytes data for manipulation
		records := readCSV(tmpFile.Name())

		t.csvRecords = records

		tpl.ExecuteTemplate(w, "display_records.gohtml", records)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	colIndex, err := strconv.Atoi(r.Form.Get("colIndex"))
	if err != nil {
		log.Fatal("Wrong value type received from validation. Check code - ", err)
	}

	fmt.Printf("User input column index: %d\n==> ", colIndex)

	tplDot := generatorPublic(colIndex)
	fmt.Printf("Successfully generated %v. \n==> ", tplDot.FilePath)
	tpl.ExecuteTemplate(w, "success.gohtml", tplDot)
}

// func generatorPublic(w http.ResponseWriter, r *http.Request) {
func generatorPublic(colIndex int) fileDelivery {
	t.colIndex = colIndex
	filename, filedir := generatorServer(t)

	tplDot := fileDelivery{
		FileName: filename,
		FileDir:  filedir,
		FilePath: filepath.Join(filedir, filename),
	}

	return tplDot
}

func generatorServer(t csvData) (string, string) {
	filenames := sortItOut(t.csvRecords, t.colIndex)
	filename, filedir := zippyZip(filenames, "marksheets.zip")
	return filename, filedir
}
