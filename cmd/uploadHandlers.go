package main

import (
	"fmt"
	mime "github.com/gabriel-vasile/mimetype"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// uploadSingle processes the uploaded file and creates a temp file for working. uploadSingle returns the temp file as *os.File , and the original filename of the uploaded file as a string.
func uploadSingle(w http.ResponseWriter, r *http.Request) (*os.File, string) {
	// Check if "tmp" dir exists; if not, make tmp dir
	if _, err := os.Stat("tmp"); os.IsNotExist(err) {
		err := os.Mkdir("tmp", 0755)
		if err != nil {
			log.Fatal("Unable to create tmp folder - ", err)
		}
	}

	// Get file from http request
	// r.ParseMultipartForm(32 << 20)

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

	return tmpFile, header.Filename
}

// uploadMultiple processes multiple uploaded files and creates temp files for each of them for working. uploadMultiple returns a slice of the temp files as []*os.File.
func uploadMultiple(w http.ResponseWriter, r *http.Request) []*os.File {
	var xfiles []*os.File

	// Check if "tmp" dir exists; if not, make tmp dir
	if _, err := os.Stat("tmp"); os.IsNotExist(err) {
		err := os.Mkdir("tmp", 0755)
		if err != nil {
			log.Fatal("Unable to create tmp folder - ", err)
		}
	}

	// Get file from http request
	r.ParseMultipartForm(32 << 20)

	files := r.MultipartForm.File["userFiles"]

	// Range over multiple files
	for _, fileHeader := range files {
		// Open each file
		file, err := fileHeader.Open()
		if err != nil {
			log.Fatal("Unable to open ", fileHeader.Filename)
		}
		defer file.Close()

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
			log.Fatalf("Uploaded file %v of invalid file format.\n", fileHeader.Filename)
		}

		// Print out metadata
		fmt.Printf("Uploaded File: %v; File Size: %v\n==> ", fileHeader.Filename, fileHeader.Size)

		// Create tmpFile for working
		tmpFile, err := ioutil.TempFile("tmp", fileHeader.Filename)
		if err != nil {
			log.Fatal("Unable to create temp file from upload - ", err)
		}
		defer tmpFile.Close()

		_, err = tmpFile.Write(buf)
		if err != nil {
			log.Fatalf("Unable to write tmpfile for %v - %v\n", fileHeader.Filename, err)
		}
		xfiles = append(xfiles, tmpFile)
	}
	return xfiles
}
