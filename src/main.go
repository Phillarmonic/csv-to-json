package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Define flags
	stdoutFlag := flag.Bool("stdout", false, "Output to stdout")
	outputFilename := flag.String("output", "data.json", "Output filename")
	flag.Parse()

	// Check if filename is provided
	if len(flag.Args()) < 1 {
		log.Fatalf(`Please provide a filename as an argument. Here's how you can do it:

Run the program with the CSV file name as an argument. For example, if your CSV file is named 'data.csv', you would run:

   csv-to-json data.csv

This will convert the 'data.csv' file into a JSON file named 'data.json' by default.

If you want to output the result to stdout instead of a file, use the '-stdout' flag:

   csv-to-json -stdout data.csv

If you want to specify a different output file name, use the '-output' flag:

   csv-to-json -output output.json data.csv

This will convert the 'data.csv' file into a JSON file named 'output.json'.`)
	}

	// Open the CSV file
	csvFile, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {
			log.Fatalf("Failed to close CSV file: %v", err)
		}
	}(csvFile)

	// Read the CSV file
	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV file: %v", err)
	}

	// Convert CSV data to a slice of maps
	var jsonData []map[string]string
	headers := records[0]
	for _, row := range records[1:] {
		item := make(map[string]string)
		for i, value := range row {
			item[headers[i]] = value
		}
		jsonData = append(jsonData, item)
	}

	// Convert to JSON
	jsonBytes, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		log.Fatalf("Failed to convert to JSON: %v", err)
	}

	// Output to stdout or write to file
	if *stdoutFlag {
		fmt.Println(string(jsonBytes))
	} else {
		err = os.WriteFile(*outputFilename, jsonBytes, 0644)
		if err != nil {
			log.Fatalf("Failed to write JSON file: %v", err)
		}
		fmt.Printf("CSV data successfully converted to JSON and saved to %s\n", *outputFilename)
	}
}
