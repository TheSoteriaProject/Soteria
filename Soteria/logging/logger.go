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

func StoreJsonLogs(log_data []byte) {
	err := os.WriteFile("../../logs/log.json", log_data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func JsonLogger(LineNumber int, Variable string, Definition string, ErrorType string) {
	log := &Log{LineNumber: LineNumber, Variable: Variable, Definition: Definition, ErrorType: ErrorType}
	printLog, err := json.MarshalIndent(log, "", "\t")
	StoreJsonLogs(printLog)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", printLog)
}
