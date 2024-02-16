package logger

import (
	"encoding/json"
	"fmt"
	"os"
)

type Log struct {
	LineNumber int
	Variable   string
	Definition string
	ErrorType  string
}

func StoreJsonLogs(log_data Log) {
	filename := "../../logs/log.json"

	// Check if File exist if not Create
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Read existing JSON data from file
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Initialize log slice
	var logs []Log

	// Check if JSON data is empty
	if len(jsonData) > 0 {
		// Unmarshal JSON data into a slice of structs if it's an array
		if err := json.Unmarshal(jsonData, &logs); err != nil {
			// If unmarshalling fails, try unmarshalling as a single object
			var log Log
			if err := json.Unmarshal(jsonData, &log); err != nil {
				fmt.Println("Error unmarshalling JSON:", err)
				return
			}
			// Append the single object to the logs slice
			logs = append(logs, log)
		}
	}

	// Append new data
	logs = append(logs, log_data)

	// Marshal the updated data back into JSON format
	updatedJsonData, err := json.MarshalIndent(logs, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Write the updated JSON data back to the file
	err = os.WriteFile(filename, updatedJsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func JsonLogger(LineNumber int, Variable string, Definition string, ErrorType string) {
	log := &Log{LineNumber: LineNumber, Variable: Variable, Definition: Definition, ErrorType: ErrorType}
	printLog, err := json.MarshalIndent(log, "", "\t")
	StoreJsonLogs(*log)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", printLog)
}

func DestroyJsonLog() error {
	filename := "../../logs/log.json"
	if err := os.Truncate(filename, 0); err != nil {
		return err
	}
	if err := os.Remove(filename); err != nil {
		return err
	}
	return nil
}
