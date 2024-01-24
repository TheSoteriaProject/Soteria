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

	err = err_given;
	if err != nil {
		log.Println("Error:", err)
	}
}

func help() {
	fmt.Println("Soteria is a linter that is used for detecting insecure communication.");
	fmt.Println("\nFlags:");
	fmt.Println("\t --Warn");
}

func version() {
	fmt.Println("Version v0.0.1");
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
	warn := false;
	flag := "nil";

	// Need to confirm it is a file path?
	if len(os.Args) > 1 {
		arg := os.Args[1];
		if len(os.Args) <= 2 {
			_ = flag; // Prevents Not Used Error
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
				_ = warn; // Prevents Not Used Error
				help();
			case "--version", "version":
				_ = warn; // Prevents Not Used Error
				version();
			default:
				run(warn);
		}
	} else {
		_ = warn; // Prevents Not Used Error
		_ = flag; // Prevents Not Used Error
		fmt.Println("No Input given try --help.");
		err := errors.New("No Input Given.");
		logError(err);
	}
}
