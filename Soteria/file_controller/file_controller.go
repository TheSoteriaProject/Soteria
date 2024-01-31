package file_controller

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"
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
func GetAllFilesAndFolders(path string, info os.FileInfo, err error) ([]string, []string) {
	files := []string{}
        folders := []string{} 
	
	if err != nil {
		fmt.Println("Error:", err)
		return files, folders
	}

	if info.IsDir() {
		// fmt.Printf("Directory: %s\n", path)
		folders = append(folders, path) // Add Token?
	} else {
		// fmt.Printf("File: %s\n", path)
		files = append(files, path) // Add Token?
	}
	return files, folders
}

func FilterFileExtensions(files []string) {
	for _, files := range files {
		split := strings.Split(files, "/")
		extension := filepath.Ext(split[len(split)-1])
		// first checks extension second checks filename and if equal to Makefile, Dockerfile, etc..
		// Makefile Check
		if strings.Contains(strings.ToLower(extension), strings.ToLower(".makefile")) || strings.Contains(strings.ToLower(split[len(split)-1]), strings.ToLower("makefile")) {
			fmt.Println(split[len(split)-1])
		}
		// Dockerfile Check
		if strings.Contains(strings.ToLower(extension), strings.ToLower(".dockerfile")) || strings.Contains(strings.ToLower(split[len(split)-1]), strings.ToLower("dockerfile")) {
			fmt.Println(split[len(split)-1])
		}
		// Bash Check
		if strings.Contains(strings.ToLower(extension), strings.ToLower(".sh")) {
			fmt.Println(split[len(split)-1])
		}
	}
}

// Main Controller For File Controller
func FileController(path string) {
	// Gather Files and Folders
	files, folders := make([]string, 0), make([]string, 0)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		file, folder := GetAllFilesAndFolders(path, info, err)
		files = append(files, file...)
		folders = append(folders, folder...)
		return nil
	})
	if err != nil {
		fmt.Println("Error walking the path:", err)
		return
	}
	

	// Extract bash, Makefiles, and Dockerfiles
	FilterFileExtensions(files)
	
	// Next
}
