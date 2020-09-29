package pkg

import (
	"encoding/csv"
	"fmt"
	"os"
)

func CsvSave(appendEmail []string, count int) {
	csvFilename := fmt.Sprintf("email_%v.csv", count)
	file, err := os.OpenFile(csvFilename, os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	csvWriter := csv.NewWriter(file)

	for _, v := range appendEmail {
		x := []string{v}
		strWrite := [][]string{x}
		csvWriter.WriteAll(strWrite)

	}
	csvWriter.Flush()
	appendEmail = nil
	return
}
