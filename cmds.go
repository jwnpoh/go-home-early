package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
)

// generator takes a csv file and generates separate csv files sorted according to user-defined criteria (e.g. sort by tutor name)
func generatorServer(t templateData) []string {
	filenames := sortItOut(t.csvRecords, t.colIndex)
	return filenames
}

func readCSV(f string) [][]string {
	// Open file to read bytes
	file, err := os.Open(f)
	if err != nil {
		log.Fatal("error opening template: ", err)
	}
	defer file.Close()

	// Encode bytes to *csv.Reader
	r := csv.NewReader(file)
	r.FieldsPerRecord = -1 // Disable checking for correct number of fields to allow for variable number of fields per record

	recs := make([][]string, 0, 50)
	for {
		rec, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("cannot read template records ", err)
		}
		recs = append(recs, rec)
	}
	return recs
}

func writeCSV(recs [][]string, fileName string) (string, error) {
	// Prepare output file
	out, err := os.Create(fileName)
	w := csv.NewWriter(out)
	defer out.Close()

	for _, k := range recs {
		err = w.Write(k)
		if err != nil {
			return out.Name(), err
		}
		w.Flush()
	}
	return out.Name(), nil
}

func sortItOut(recs [][]string, colIndex int) []string {
	// Prepare mapping data structures
	tutorsMap := make(map[string][][]string)
	outRecs := make([][]string, 0, 25)

	// Do mapping of student recs by tutor
	for n, i := range recs {
		if n < 1 {
			outRecs = append(outRecs, i)
			continue
		}
		tutorsMap[i[colIndex]] = append(tutorsMap[i[colIndex]], i)
	}

	// Prepare output folder
	if _, err := os.Stat(filepath.Join("tmp", "marksheets")); os.IsNotExist(err) {
		err := os.Mkdir(filepath.Join("tmp", "marksheets"), 0755)
		if err != nil {
			log.Fatal("Unable to create output folder: ", err)
		}
	}

	var xOutName []string
	// Prepare writing to file
	for i, j := range tutorsMap {
		k := make([][]string, 0, 25)

		// Insert header row to be written before student recs
		k = append(k, outRecs...)

		// Insert student recs
		k = append(k, j...)

		// Write CSV files
		fileName := filepath.Join("tmp", "marksheets", i) + ".csv"
		outName, err := writeCSV(k, fileName)
		if err != nil {
			os.Remove(outName)
			log.Fatal("Parsing unsuccessful. Did you enter the correct column index?")
			return nil
		}
		xOutName = append(xOutName, outName)
	}
	return xOutName
}
