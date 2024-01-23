package main;

import (
	"fmt"
	"os"
	// "os/exec"
	"errors"
	"log"
)

func simulateError() error {
	return fmt.Errorf("This is a simulated error")
}

func logError(err_given error) {
	logFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	
	// Change so you can pass in function???
	err = err_given;
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

func run(flag bool) {
	fmt.Println("Running...");
	error_flag := false;

	if flag == true {
		fmt.Println("Warning Mode: On");
	} else {
		fmt.Println("Warning Mode: Off");
	}

	// probably not needed here but will be used on each function
	if error_flag == true {
		err := errors.New("Fake Error for now.");
		logError(err);
	}
}

func main() {
	// Variables
	warn := false;
	flag := "nil";

	if len(os.Args) > 1 {
		arg := os.Args[1];
		if len(os.Args) <= 2 {
			_ = flag;
			warn = false;
		} else {
			flag = os.Args[2];
			switch flag {
				case "--warn", "warn":
					warn = true;
				default:
					warn = false;
			}
		}

		switch arg {
			case "--help", "help":
				_ = warn; // Not Used Error Prevent
				help();
			case "--version", "version":
				_ = warn; // Not Used Error Prevent
				version();
			default:
				run(warn);
		}
	} else {
		_ = warn; // Not Used Error Prevent
		_ = flag;
		fmt.Println("No Input given try --help.");
		err := errors.New("No Input Give");
		logError(err);
	}
}
