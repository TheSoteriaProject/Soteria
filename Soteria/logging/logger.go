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

func StoreJsonLogs(log_data Log) {
	file, err := os.OpenFile("../logs/log.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	existingData, err := os.ReadFile("../logs/log.json")
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

	if err := os.WriteFile("../logs/log.json", jsonData, 0644); err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}

func JsonLogger(FileName string, LineNumber int, Line string, Issue string, ErrorType string) {
	log := &Log{FileName: FileName, LineNumber: LineNumber, Line: Line, Issue: Issue, ErrorType: ErrorType}
	printLog, err := json.MarshalIndent(log, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling log:", err)
		return
	}
	fmt.Printf("%s\n", printLog)

	// Store log data
	StoreJsonLogs(*log)
}

// Not Implemented
func DestroyJsonLog() error {
	filename := "../logs/log.json"
	if err := os.Truncate(filename, 0); err != nil {
		return err
	}
	if err := os.Remove(filename); err != nil {
		return err
	}
	return nil
}
