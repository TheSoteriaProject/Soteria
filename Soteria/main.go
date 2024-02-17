package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"

	// Custom Files
	"Soteria/diverter"
	"Soteria/file_controller"
)

func main() {
	fmt.Println("Welcome to Insecure Communication Linter.")

	// Confrim/Test File Controller Connection
	file_controller.TestConnection()

	//Confirm/Test Diverter Connection
	diverter.TestConnection()

	if len(os.Args) > 1 {
		// Take Project Path
		input := os.Args[1]

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

		// Parse Flag
		flag.Parse()

		if runTests {
			cmd := exec.Command("go", "test", "./...", "-v")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			fmt.Println("Running tests...")
			err := cmd.Run()
			if err != nil {
				fmt.Printf("Error running tests: %v\n", err)
				os.Exit(1)
			}
		} else if helpUser {
			// Create A Help Controller
			fmt.Println("Help Page")
		} else if versionCheck {
			// Set From File
			version := "0.0.0.0"
			fmt.Println("Version: v" + version)
		} else {
			// Add Other flags
			// --warn etc...
			// Check for multiple flags possibly

			// If File Path Does Exist
			if _, err := os.Stat(input); err == nil {
				// All the Files that are to be checked.
				file_pool := file_controller.FileController(input)
				// Divert Files to correct parser || parsers
				_, err := diverter.DivertFiles(file_pool)
				if err != nil {
					fmt.Printf("Error Running Diverter: %v\n", err)
					os.Exit(1)
				}
			} else if errors.Is(err, os.ErrNotExist) {
				// If Path Does Not Exist Throw Error and Exit
				fmt.Println("Path Does Not Exist.")
				os.Exit(1)
			} else {
				// If File Does Not Exist Throw Err and EXIT
				// SHOULD NOT MAKE IT HERE
				fmt.Println("File Does Not Exist")
				os.Exit(1)
			}
		}
	} else {
		// Invalid Input or Bad CLI Argument
		fmt.Println("It seems you have given an invalid input. Try --help")
	}
}
