package logger

import (
	"encoding/json"
	"fmt"
	"os"
)

type Log struct {
	FileName   string
	LineNumber int
	Line       string
	Issue      string
	ErrorType  string
}

// StoreJsonLogs takes in the JSON data and writes it to a file so it can be read.
func StoreJsonLogs(log_data Log) {
	file, err := os.OpenFile("../logs/bash_log.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	existingData, err := os.ReadFile("../logs/bash_log.json")
	if err != nil {
		fmt.Println("Error reading JSON from file:", err)
		return
	}

	if len(existingData) > 0 {
		existingData = existingData[:len(existingData)-1]
	}

	if len(existingData) == 0 {
		existingData = append(existingData, '[')
	}

	if len(existingData) > 1 {
		existingData = append(existingData, ',')
	}

	jsonData, err := json.Marshal(log_data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	jsonData = append(existingData, jsonData...)
	jsonData = append(jsonData, ']')

	if err := os.WriteFile("../logs/bash_log.json", jsonData, 0644); err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}

// Json Logger handles json in take, prints,  and sends it off to be store in file.
func JsonLogger(FileName string, LineNumber int, Line string, Issue string, ErrorType string, DisableLogPrint bool) {
	log := &Log{FileName: FileName, LineNumber: LineNumber, Line: Line, Issue: Issue, ErrorType: ErrorType}
	printLog, err := json.MarshalIndent(log, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling log:", err)
		return
	}

	if !DisableLogPrint {
		fmt.Printf("%s\n", printLog)
	}

	// Store log data
	StoreJsonLogs(*log)
}

// DestroyJsonLog truncates already made file and deletes log(s)
func DestroyJsonLog() error {
	filename := "../logs/bash_log.json"
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		fmt.Printf("File '%s' does not exist\n", filename)

		if err := os.Truncate(filename, 0); err != nil {
			return err
		}
		if err := os.Remove(filename); err != nil {
			return err
		}
	}
	return nil
}
