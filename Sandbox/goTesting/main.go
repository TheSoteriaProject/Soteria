package main;

import (
	"fmt"
	"os"
	"log"
)

func simulateError() error {
	return fmt.Errorf("This is a simulated error")
}

func logError() {
	logFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	
	// Change so you can pass in function???
	err = simulateError()
	if err != nil {
		log.Println("Error:", err)
	}
}

func help() {
	fmt.Println("Soteria is a linter that is used for...");
	fmt.Println("\nFlags:");
	fmt.Println("\t --Warn");
}

func version() {
	fmt.Println("Version 0.0.1");
}

func run(flag string) {
	fmt.Println("Running...");
	
	// warning_flag := false;
	error_flag := true;

	if flag == "--warn" {
		// warning_flag = true;
		fmt.Println("Warning Mode: On");
	} else {
		fmt.Println("Warning Mode: Off");
	}

	// probably not needed here but will be used on each function
	if error_flag == true {
		logError();
	}
}

func main() {
	if len(os.Args) > 1{
		arg := os.Args[1];
		switch arg {
			case "--help", "help":
				help();
			case "--version", "version":
				version();
			case "--warn":
				run(arg);
			default:
				fmt.Println("Invalid CLI argument.\nTry --help");
		}
	} else {
		run("NULL");
	}
}
