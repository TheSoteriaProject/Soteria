package main

import (
	// Standard Packages.
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	// Custom Packages(Sorta Packages lol)
	"Soteria/diverter"
	"Soteria/file_controller"
	Help "Soteria/help"       // Alias Given
	JLogger "Soteria/logging" // Alias Given
)

func main() {
	// Intro Print NOT neeed but is nice.
	fmt.Println("Welcome to Insecure Communication Linter.")

	/*****************************
	// Down Below is old Test Connections. Not Needed but fine to keep.

	// Confrim/Test File Controller Connection
	// file_controller.TestConnection()

	//Confirm/Test Diverter Connection
	// diverter.TestConnection()
	*****************************/

	// Checks to see if there was an argument pass to the command line.
	if flag.NFlag() >= 0 || flag.NArg() >= 0 {
		// Diasble Terminal Flags
		flag.Usage = func() {}

		// Test Flag
		var runTests bool
		flag.BoolVar(&runTests, "test", false, "Run tests")

		// Help Flag
		var helpUser bool
		flag.BoolVar(&helpUser, "help", false, "Help User")

		// Version Flag
		var versionCheck bool
		flag.BoolVar(&versionCheck, "version", false, "Version Check")
		flag.BoolVar(&versionCheck, "v", false, "Version Check")

		// Bash, Makefile, and Docker Flag
		enableBash := flag.Bool("enableBash", true, "Check Bash Files")
		enableMakefile := flag.Bool("enableMakefile", true, "Check Makefiles File")
		enableDockerfile := flag.Bool("enableDockerfile", true, "Check Dockerfile Files")

		// Logging Print Flag
		enableLogPrint := flag.Bool("enableLogPrint", true, "Check If Logs Print")

		// Warn Flag
		var warnUser bool
		flag.BoolVar(&warnUser, "warn", false, "Warn User")

		// Parse Flag
		flag.Parse()

		// Take Project Path and set it.
		args := flag.Args()
		input := ""
		if len(args) > 0 {
			input = args[len(args)-1]
		} else {
			input = flag.Arg(0)
		}

		if runTests {
			// If the Unit Test flag is passed it will run the test and confirm any changes and current tool work ho it should.
			cmd := exec.Command("go", "test", "./...", "-v")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			// Let the user know.
			fmt.Println("Running tests...")
			err := cmd.Run()
			if err != nil {
				fmt.Printf("Error running tests: %v\n", err)
				os.Exit(1)
			}
		} else if helpUser {
			// If the help flag is passed provide the the help page.
			Help.Info()
			os.Exit(0)
		} else if versionCheck {
			// If the version flag is passed provide the version for debug and relevance reasons.
			// This gets the current tag. More complex and stupid than it should be, but is how a stack overflow post does it.
			tagOutput := exec.Command("git", "rev-list", "--tags", "--max-count=1")
			outTag, errTag := tagOutput.Output()
			if errTag != nil {
				fmt.Println("Error: ", errTag)
			}

			tag := strings.TrimSpace(string(outTag))

			versionCmd := exec.Command("git", "describe", "--tags", tag)
			outVersion, errVersion := versionCmd.Output()
			if errVersion != nil {
				fmt.Println("Error:", errVersion)
				return
			}

			fmt.Println("Version:", strings.TrimSpace(string(outVersion)))
			os.Exit(0)
		} else {
			// Main Sequence
			// If File Path Does Exist
			if _, err := os.Stat(input); err == nil {
				// All the Files that are to be checked.
				file_pool := file_controller.FileController(input, *enableMakefile, *enableDockerfile, *enableBash)

				destroy_logs := []string{"../logs/bash_log.json", "../logs/dockerfile_log.json", "../logs/makefile_log.json"} // Should be dynamic based on logs generated but not enough time.
				for _, logfile := range destroy_logs {
					err := JLogger.DestroyJsonLog(logfile) // Truncates Old File before start if not possible error. Lovely...
					if err != nil {
						fmt.Println("Error Destroying Logs.")
						os.Exit(1)
					}
				}

				// Take the file pool and divert to each analyzer.
				diverter.DivertFiles(file_pool, warnUser, *enableMakefile, *enableDockerfile, *enableBash, *enableLogPrint)

				// Files in bad format and Warn Flag won't work on them becuase the status is determined indiidually instead of being
				// being determined either by line or by choice. Both are good just mis-communication issues.
				file_logs := []string{"../logs/bash_log.json"} //, "../logs/dockerfile_log.json", "../logs/makefile_log.json"} // Should be dynamic based on logs generated but not enough time.
				status := 0
				r_status := 0
				for _, filename := range file_logs {
					status = JLogger.CheckForReturnType(filename)
					// If status fails at any point it stays 1 so it exits with error.
					if status != 0 {
						r_status = 1
					}
				}
				os.Exit(r_status)
			} else if errors.Is(err, os.ErrNotExist) {
				// If Path Does Not Exist Throw Error and Exit
				fmt.Println("It seems you have given an invalid input. Try --help")
				os.Exit(1)
			} else {
				// If File Does Not Exist Throw Err and EXIT
				fmt.Println("File Does Not Exist") // Shoud not technically make it here but in-case.
				os.Exit(1)
			}
		}
	} else {
		// No Input Flag is given or bad Command Line Input.
		fmt.Println("It seems you have given an invalid input. Try --help")
		os.Exit(1)
	}
}
