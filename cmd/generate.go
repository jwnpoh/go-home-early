package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

var t CsvData

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
		tmpFile, _ := uploadSingle(w, r)

		// Take file bytes data for manipulation
		records := readCSV(tmpFile.Name())

		t.CsvRecords = records
		t.FunctionPath = "/generate/upload"

		tpl.ExecuteTemplate(w, "display_records.gohtml", t)
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

func generatorPublic(colIndex int) FileDelivery {
	t.ColIndex = colIndex
	filename, filedir := generatorServer(t)

	tplDot := FileDelivery{
		FileName: filename,
		FileDir:  filedir,
		FilePath: filepath.Join(filedir, filename),
	}

	return tplDot
}

func generatorServer(t CsvData) (string, string) {
	filenames := sortItOut(t.CsvRecords, t.ColIndex)
	filename, filedir := zippyZip(filenames, "marksheets.zip")
	return filename, filedir
}
