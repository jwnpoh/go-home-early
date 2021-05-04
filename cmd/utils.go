package main

import (
	"archive/zip"
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Information struct {
	Title       string
	Path        string
	Description string
}

type CsvData struct {
	CsvRecords   [][]string
	ColIndex     int
	FunctionPath string
}

type FileDelivery struct {
	FileName string
	FileDir  string
	FilePath string
}

// readCSV opens a csv file and reads it into a [][]string
func readCSV(f string) [][]string {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal("error opening template: ", err)
	}
	defer file.Close()

	defer os.Remove(file.Name())

	r := csv.NewReader(file)
	r.FieldsPerRecord = -1

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

// writeCSV takes a [][]string and writes csv encoded file specified by the given filename and returns the filename of the written file and an error if write is not successful.
func writeCSV(recs [][]string, fileName string) (string, error) {
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

// sortItOut implements a simple sorting algorithm to generate templates according to user-defined criteria
func sortItOut(recs [][]string, colIndex int) []string {

	tutorsMap := make(map[string][][]string)
	outRecs := make([][]string, 0, 25)

	for n, i := range recs {
		if n < 1 {
			outRecs = append(outRecs, i)
			continue
		}
		tutorsMap[i[colIndex]] = append(tutorsMap[i[colIndex]], i)
	}

	if _, err := os.Stat(filepath.Join("tmp", "marksheets")); os.IsNotExist(err) {
		err := os.Mkdir(filepath.Join("tmp", "marksheets"), 0755)
		if err != nil {
			log.Fatal("Unable to create output folder: ", err)
		}
	}

	var xOutName []string
	for i, j := range tutorsMap {
		k := make([][]string, 0, 25)

		k = append(k, outRecs...)
		k = append(k, j...)

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

func zippyZip(files []string, filename string) (string, string) {
	newZipFile, err := os.Create(filepath.Join("tmp", filename))
	if err != nil {
		log.Fatalf("Unable to create new archive %s - %v\n", filename, err)
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	for _, file := range files {
		if err = addFileToZip(zipWriter, file); err != nil {
			log.Fatalf("Unable to add file %s to archive - %v\n", file, err)
		}
	}
	return filename, "tmp"
}

func addFileToZip(zipWriter *zip.Writer, filename string) error {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = filepath.Base(fileToZip.Name())
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, fileToZip)
	return err
}
