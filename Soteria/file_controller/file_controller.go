package file_controller

import (
	"fmt"
	"os"
)

// Confirm File Connection
func TestConnection() {
	fmt.Println("File Controller: Active")
}

// Show Files and Folders
func ShowSliceData(path []string) {
	for _, path := range path {
		fmt.Printf("%s\n", path)
	}
}

// Path Traversal
func GetAllFilesAndFolders(path string, info os.FileInfo, err error) error {
	files := []string{}
	folders := []string{}

	if err != nil {
		fmt.Println(err)
		return nil
	}

	if info.IsDir() {
		// fmt.Printf("Directory: %s\n", path)
		folders = append(folders, path + " : folder")
	} else {
		// fmt.Printf("File: %s\n", path)
		files = append(files, path + " : file")
	}
	
	// Adjust Naming
	ShowSliceData(files)
	ShowSliceData(folders)

	return nil
}


func FileController() {
	// Controls all the functions
}
