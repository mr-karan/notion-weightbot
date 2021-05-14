package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func saveToCSV(date string, weight float64, path string) {
	var data [][]string
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Append the record in CSV.
	data = append(data, []string{date, fmt.Sprintf("%f", weight)})
	w := csv.NewWriter(file)
	w.WriteAll(data)
}
