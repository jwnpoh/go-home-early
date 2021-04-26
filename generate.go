package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	mime "github.com/gabriel-vasile/mimetype"
)

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
	file, header, err := r.FormFile("userFile")
	if err != nil {
		log.Fatal("Error retrieving file - ", err)
	}
	defer file.Close()

	// Print out metadata
	fmt.Printf("Uploaded File: %v\n", header.Filename)
	fmt.Printf("File Size: %v\n", header.Size)
	fmt.Printf("MIME Header: %v\n", header.Header)

	buf := make([]byte, 0, 512)

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Unable to read uploaded file - ", err)
	}
	buf = append(buf, fileBytes...)

	// Check that the correct filetype is uploaded
	filetype := mime.Detect(fileBytes)
	if filetype.String() != "text/csv" {
		http.Error(w, "The provided file format is not allowed. Please upload only CSV files", http.StatusBadRequest)
		fmt.Println("Uploaded file of invalid file format.")
		return
	}
	// Create tmpFile for working
	tmpFile, err := ioutil.TempFile("tmp", "tmp*.csv")
	if err != nil {
		log.Fatal("Unable to create temp file from upload - ", err)
	}

	_, err = tmpFile.Write(buf)
	if err != nil {
		log.Fatal("Unable to write file - ", err)
	}

	tmpFile.Close()

	err = os.Rename(tmpFile.Name(), filepath.Join("tmp", header.Filename))

	if err != nil {
		log.Fatal("Error renaming file - ", err)
	}

	http.ServeFile(w, r, filepath.Join("tmp", header.Filename))
}
