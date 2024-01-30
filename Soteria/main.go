package main

import (
	"fmt"
	"errors"
	// "os/exec"
	"os"
	"path/filepath"

	// Custom Files
	"Soteria/file_controller"
)

func main() {
	fmt.Println("Welcome to Insecure Communication Linter.")
	file_controller.TestConnection()

	// Check Length
	if len(os.Args) > 1 {
		// Grab file Path from CL Arg
		file_path := os.Args[1]

		// Check if Exist
		if _, err := os.Stat(file_path); err == nil {
			filepath.Walk(file_path, file_controller.GetAllFiles)
			if err != nil {
				fmt.Printf("error gathering path %v: %v\n", file_path, err)
			}
		} else if errors.Is(err, os.ErrNotExist) {
  			// Path Does Not Exist
			// Add Error
			os.Exit(1)
		} else {
			// File Does Not Exist
			// Add Error
			os.Exit(1)
		}
	} else {
		fmt.Println("It seems you have given an invalid input. Try -- help") 
	}
}
