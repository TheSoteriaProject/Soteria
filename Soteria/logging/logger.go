package logger

import (
	"encoding/json"
	"fmt"
	"os"
)

type Log struct {
	FileName   string `json:"FileName"`
	LineNumber int    `json:"LineNumber"`
	Line       string `json:"Line"`
	Issue      string `json:"Issue"`
	ErrorType  string `json:"ErrorType"`
}

// CheckForReturnType atakes the JSON log and based on if an ErrorType of "Error" is found determines the Exit code for the program
func CheckForReturnType() int {
	filename := "../logs/bash_log.json"
	file, err := os.Open(filename) // Add log or pre-check becuase may not exist
	if err != nil {
		// fmt.Println("Failed to open JSON Logs:", err)
		os.Exit(0) // Ehhhh not certain
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var logs []Log
	if err := decoder.Decode(&logs); err != nil {
		fmt.Println("Failed to decode JSON:", err)
	}

	for _, log := range logs {
		if log.ErrorType == "Error" {
			return 1
		}
	}

	return 0
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
func JsonLogger(FileName string, LineNumber int, Line string, Issue string, ErrorType string, enableLogPrint bool) {
	log := &Log{FileName: FileName, LineNumber: LineNumber, Line: Line, Issue: Issue, ErrorType: ErrorType}
	printLog, err := json.MarshalIndent(log, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling log:", err)
		return
	}

	if enableLogPrint {
		fmt.Printf("%s\n", printLog)
	}

	// Store log data
	StoreJsonLogs(*log)
}

// DestroyJsonLog truncates already made file and deletes log(s)
func DestroyJsonLog() error {
	filename := "../logs/bash_log.json"
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		if err := os.Truncate(filename, 0); err != nil {
			return err
		}
		if err := os.Remove(filename); err != nil {
			return err
		}
	}
	return nil
}
