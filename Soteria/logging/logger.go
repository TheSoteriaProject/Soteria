package logger

import (
	"encoding/json"
	"fmt"
)

type Log struct {
	LineNumber int
	Variable   string
	Definition string
	ErrorType  string
}

func StoreJsonLogs(log *Log) {
	// Not Implemented
	// Stor Logs in File??
	fmt.Println("Stored Logs: ", log)
}

func JsonLogger(LineNumber int, Variable string, Definition string, ErrorType string) {
	log := &Log{LineNumber: LineNumber, Variable: Variable, Definition: Definition, ErrorType: ErrorType}
	StoreJsonLogs(log)
	printLog, err := json.MarshalIndent(log, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", printLog)
}
