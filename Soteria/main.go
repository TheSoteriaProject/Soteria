package main

import (
	"fmt"
	"errors"
	"os"
	"testing"

	// Custom Files
	"Soteria/file_controller"
	"Soteria/testing_controller"
)

func main() {
	fmt.Println("Welcome to Insecure Communication Linter.")
	
	// Confrim/Test File Controller Connection
	file_controller.TestConnection()

	if len(os.Args) > 1 {
		input := os.Args[1]
		if input == "--test" {
			// Add Testing Controller
			fmt.Println("Testing Tool")
			testing_controller.TestingController(&testing.M{})
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
