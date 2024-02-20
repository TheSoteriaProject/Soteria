package logger

import (
	"encoding/json"
	"fmt"
	"os"
)

type Log struct {
	FileName   string
	LineNumber int
	Variable   string
	Definition string
	ErrorType  string
}

func StoreJsonLogs(log_data Log) {
	// Not Implemnted
}

func JsonLogger(FileName string, LineNumber int, Variable string, Definition string, ErrorType string) {
	log := &Log{FileName: FileName, LineNumber: LineNumber, Variable: Variable, Definition: Definition, ErrorType: ErrorType}
	printLog, err := json.MarshalIndent(log, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling log:", err)
		return
	}
	fmt.Printf("%s\n", printLog)

	// Store log data
	// StoreJsonLogs(*log)
}

// Not Implemented
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
