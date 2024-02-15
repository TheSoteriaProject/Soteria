package logger

import (
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
	fmt.Println("Stored Logs: ", log)
}

func JsonLogger(LineNumber int, Variable string, Definition string, ErrorType string) {
	log := &Log{LineNumber: LineNumber, Variable: Variable, Definition: Definition, ErrorType: ErrorType}
	StoreJsonLogs(log)
	fmt.Printf("Line Number: %d, Variable: %s, Definition: %s, Error Type: %s\n", log.LineNumber, log.Variable, log.Definition, log.ErrorType)
}
