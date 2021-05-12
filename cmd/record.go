package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

type RecData struct {
	CsvData
	XInputRecs [][][]string
	OutputName string
  Information
}

var rec RecData

func record(w http.ResponseWriter, r *http.Request) {
	info := info["Record"]
	err := tpl.ExecuteTemplate(w, "record.gohtml", info)
	if err != nil {
		log.Fatal("unable to execute template - ", err)
	}
	fmt.Printf("Selected function: %v\n==> ", info.Title)
}

func recordUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		mainResList, origFileName := uploadSingle(w, r)
		inputfiles := uploadMultiple(w, r)

		masterRecs := readCSV(mainResList.Name())

		xinputRecs := make([][][]string, 0, 10)
		for _, l := range inputfiles {
			inFileName := l.Name()
			inRecs := readCSV(inFileName)

			// check header
			for i, k := range inRecs[0] {
				if k != masterRecs[0][i] {
					http.Error(w, fmt.Sprintf("header rows do not match for %v. Please check that you have provided the correct files", l.Name()), http.StatusBadRequest)
					log.Fatalf("Header rows of the uploaded file %v do not match with central results list %v - ", l.Name(), mainResList.Name())
				}
			}
			xinputRecs = append(xinputRecs, inRecs)
		}

		rec.OutputName = filepath.Join("tmp", origFileName)
		rec.XInputRecs = xinputRecs
		rec.CsvRecords = masterRecs
		rec.FunctionPath = "/record/upload"
    rec.Information = info["Record"]

		tpl.ExecuteTemplate(w, "display_records.gohtml", rec)
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

	rec.ColIndex = colIndex

	fmt.Printf("User input column index: %d\n==> ", colIndex)

	tplDot := recordMain(rec.CsvRecords, rec.XInputRecs, rec.ColIndex)

	fmt.Printf("Successfully generated %v. \n==> ", tplDot.FilePath)
	tpl.ExecuteTemplate(w, "success.gohtml", tplDot)
}

func recordMain(f [][]string, xf [][][]string, n int) FileDelivery {

	outRecs := make([][]string, 0, 50)

	for i, o := range f {
		if i < 1 {
			outRecs = append(outRecs, o)
			continue
		}
		for _, j := range xf {
			for _, l := range j {
				if l[n] == o[n] {
					outRecs = append(outRecs, l)
				}
			}
		}
	}

	generatedFileName, err := writeCSV(outRecs, rec.OutputName)
	if err != nil {
		log.Fatal(err)
	}

	tplDot := FileDelivery{
		FileName: filepath.Base(generatedFileName),
		FilePath: generatedFileName,
		FileDir:  "tmp",
	}

	return tplDot
}
