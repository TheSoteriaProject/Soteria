package main

import (
	"fmt"
	"errors"
	"os"

	// Custom Files
	"Soteria/file_controller"
)

func main() {
	fmt.Println("Welcome to Insecure Communication Linter.")
	
	// Confrim/Test File Controller Connection
	file_controller.TestConnection()

	if len(os.Args) > 1 {
		
		file_path := os.Args[1]

		if _, err := os.Stat(file_path); err == nil {
			file_controller.FileController(file_path)
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
