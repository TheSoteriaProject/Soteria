package file_controller

import (
	"fmt"
	"os"
)

// Confirm File Connection
func TestConnection() {
	fmt.Println("Hello From File Controller.\n")
}

// Path Traversal
func GetAllFiles(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if info.IsDir() {
		fmt.Printf("Directory: %s\n", path)
	} else {
		fmt.Printf("File: %s\n", path)
	}

	return nil
}
