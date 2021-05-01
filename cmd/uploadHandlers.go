package main

import (
	"fmt"
	mime "github.com/gabriel-vasile/mimetype"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func uploadSingle(w http.ResponseWriter, r *http.Request) *os.File {
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
	fmt.Printf("Uploaded File: %v; File Size: %v\n==> ", header.Filename, header.Size)

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Unable to read uploaded file - ", err)
	}

	// Check that the correct filetype is uploaded
	filetype := mime.Detect(fileBytes)
	if filetype.String() != "text/csv" {
		http.Error(w, "The provided file format is not allowed. Please upload only CSV files", http.StatusBadRequest)
		log.Fatal("Uploaded file of invalid file format.")
	}
	// Create tmpFile for working
	buf := make([]byte, 0, 512)
	buf = append(buf, fileBytes...)
	tmpFile, err := ioutil.TempFile("tmp", "tmp*.csv")
	if err != nil {
		log.Fatal("Unable to create temp file from upload - ", err)
	}
	defer tmpFile.Close()

	_, err = tmpFile.Write(buf)
	if err != nil {
		log.Fatal("Unable to write tempfile - ", err)
	}

	return tmpFile
}

func uploadMultiple(w http.ResponseWriter, r *http.Request) {
	var fileNames []string

	// Check if "tmp" dir exists; if not, make tmp dir
	if _, err := os.Stat("tmp"); os.IsNotExist(err) {
		err := os.Mkdir("tmp", 0755)
		if err != nil {
			log.Fatal("Unable to create tmp folder - ", err)
		}
	}

	// Get file from http request
	r.ParseMultipartForm(32 << 20)

	files := r.MultipartForm.File["userFile"]

	// Range over multiple files
	for _, fileHeader := range files {
		// Open each file
		file, err := fileHeader.Open()
		if err != nil {
			log.Fatal("Unable to open ", fileHeader.Filename)
		}
		defer file.Close()

		// Print out metadata
		fmt.Printf("Uploaded File: %v\n", fileHeader.Filename)
		fmt.Printf("File Size: %v\n", fileHeader.Size)
		buf := make([]byte, 0, 512)

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal("Unable to read uploaded file - ", err)
		}

		buf = append(buf, fileBytes...)

		// Check that the correct filetype is uploaded
		filetype := mime.Detect(buf)
		if filetype.String() != "text/csv" {
			http.Error(w, "The provided file format is not allowed. Please upload only CSV files", http.StatusBadRequest)
			log.Fatal("Uploaded of invalid file format.")
		}

		// Create tmpFile for working
		tmpFile, err := ioutil.TempFile("tmp", "tmp*.csv")
		if err != nil {
			log.Fatal("Unable to create temp file from upload - ", err)
		}
		defer tmpFile.Close()

		tmpFile.Write(buf)

		fileNames = append(fileNames, tmpFile.Name())
	}
	fmt.Println(fileNames)
}
