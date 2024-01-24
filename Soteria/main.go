package main

import (
	"fmt"
	"errors"
	"os/exec"
	"os"
)

// File Controller to deal with sorting files
func file_controller(filePath string) {
	// Exact Path to file controller
	execPath, err := exec.LookPath("./SoteriaIgnore/file_controller")

        if err != nil {
                fmt.Println("Error: ", err)
        } else {
		// Run Rust Executable
                cmd := exec.Command(execPath, filePath)
		
		// Show Output
                cmd.Stdout = os.Stdout
                cmd.Stderr = os.Stderr

                // Run & Check for Error
		cmdErr := cmd.Run()
                if cmdErr != nil {
                        panic(cmdErr)
                }
        }
}

func main() {
	fmt.Println("Welcome to Insecure Communication Linter.")
	
	// Check Length
	if len(os.Args) > 1 {
		// Grab file Path from CL Arg
		filePath := os.Args[1]

		// Check if Exist
		if _, err := os.Stat(filePath); err == nil {
			file_controller(filePath) // Gathers Files	
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
