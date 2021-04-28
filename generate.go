package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	mime "github.com/gabriel-vasile/mimetype"
)

type templateData struct {
	csvRecords [][]string
	colIndex   int
}

var t templateData

func generate(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "generate.gohtml", cmds[0])
	if err != nil {
		log.Fatal("unable to execute template - ", err)
	}
}

func generateCmd(w http.ResponseWriter, r *http.Request) {
	// Check if "tmp" dir exists; if not, make tmp dir
	if _, err := os.Stat("tmp"); os.IsNotExist(err) {
		err := os.Mkdir("tmp", 0755)
		if err != nil {
			log.Fatal("Unable to create tmp folder - ", err)
		}
	}

	// Get file from http request
	r.ParseMultipartForm(32 << 20)

	// Gets the file, and fileheader
	file, _, err := r.FormFile("userFile")
	if err != nil {
		log.Fatal("Error retrieving file - ", err)
	}
	defer file.Close()

	// Read the file into useable []bytes
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Unable to read uploaded file - ", err)
	}

	// Check that the correct filetype is uploaded
	filetype := mime.Detect(fileBytes)
	if filetype.String() != "text/csv" {
		http.Error(w, "The provided file format is not allowed. Please upload only CSV files", http.StatusBadRequest)
		fmt.Println("Uploaded file of invalid file format.")
		return
	}

	// Create and write temp file for working
	buf := make([]byte, 0, 512)
	buf = append(buf, fileBytes...)

	tmpFile, err := ioutil.TempFile("tmp", "tmp*.csv")
	if err != nil {
		log.Fatal("Unable to create temp file from upload - ", err)
	}

	_, err = tmpFile.Write(buf)
	if err != nil {
		log.Fatal("Unable to write file - ", err)
	}

	// Take file bytes data for manipulation
	records := readCSV(tmpFile.Name())

	t.csvRecords = records

	tpl.ExecuteTemplate(w, "display_records.gohtml", records)
}

func generatorPublic(w http.ResponseWriter, r *http.Request) {
	colIndex, err := strconv.Atoi(r.FormValue("colIndex"))
	if err != nil {
		log.Fatal("Wrong value type received from validation. Check code - ", err)
	}

	t.colIndex = colIndex

	archive := generatorServer(t)

	w.Header().Set("Content-Disposition", "attachement; filename=marksheets.zip")
	http.ServeFile(w, r, archive)
}

// Serve file to user
// http.ServeFile(w, r, filepath.Join("tmp", header.Filename)) */
