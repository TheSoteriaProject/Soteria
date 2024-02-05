package main

import (
	"fmt"
	"flag"
	"errors"
	"os"
	"os/exec"
	// "testing"

	// Custom Files
	"Soteria/file_controller"
)

/func main() {
	// fmt.Println("Welcome to Insecure Communication Linter.")

	// Confrim/Test File Controller Connection
	file_controller.TestConnection()

	if len(os.Args) > 1 {
		// Main
		input := os.Args[1]
		
		// Sub
		var runTests bool
		flag.BoolVar(&runTests, "test", false, "Run tests")
		flag.Parse()

		if runTests {
			// Add Testing Controller
			// fmt.Println("Testing Tool")
			// Run tests using the "go test" command

			// Adjust and create a function that calls them all
			cmd := exec.Command("go", "test", "./...", "-v")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			fmt.Println("Running tests...")
			err := cmd.Run()
		if err != nil {
			fmt.Printf("Error running tests: %v\n", err)
			os.Exit(1)
		}
		} else if input == "--help" {
			// Create A Help Controller
			fmt.Println("Help Page")
		} else if input == "--version" || input == "--v" {
			// Set From File
			version := "0.0.0.0"
			fmt.Println("Version: v" + version)
		} else {
			// Add Other flags
			// --warn etc... 
			// Check for multiple flags possibly


			// check if File Path
			if _, err := os.Stat(input); err == nil {
				// Adjust to give a return
				file_controller.FileController(input)
			} else if errors.Is(err, os.ErrNotExist) {
  				// Path Does Not Exist
				fmt.Println("Path Does Not Exist.")
				// Add Error
				os.Exit(1)
			} else {
				// File Does Not Exist
				// SHOULD NOT MAKE IT HERE
				fmt.Println("File Does Not Exist")
				// Add Error
				os.Exit(1)
			}
		}
	} else {
		fmt.Println("It seems you have given an invalid input. Try --help") 
	}
}
