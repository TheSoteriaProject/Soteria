package file_controller

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"

	// Other
	"Soteria/ignore_file_parser"
)

// Confirm File Connection
func TestConnection() {
	fmt.Println("Testing File Controller Connection.")
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

func FilterFileExtensions(files []string, u_makefile bool, u_dockerfile bool, u_bash bool) []string {
	filtered_files := []string{}

	for _, files := range files {
		split := strings.Split(files, "/")
		extension := filepath.Ext(split[len(split)-1])
		// first checks extension second checks filename and if equal to Makefile, Dockerfile, etc..
		// Makefile Check
		if u_makefile == true && strings.Contains(strings.ToLower(extension), strings.ToLower(".makefile")) || strings.Contains(strings.ToLower(split[len(split)-1]), strings.ToLower("makefile")) {
			// fmt.Println(split[len(split)-1])
			filtered_files = append(filtered_files, split[len(split)-1])
		}
		// Dockerfile Check
		if u_dockerfile == true && strings.Contains(strings.ToLower(extension), strings.ToLower(".dockerfile")) || strings.Contains(strings.ToLower(split[len(split)-1]), strings.ToLower("dockerfile")) {
			// fmt.Println(split[len(split)-1])
			filtered_files = append(filtered_files, split[len(split)-1])
		}
		// Bash Check
		if u_bash == true && strings.Contains(strings.ToLower(extension), strings.ToLower(".sh")) {
			// fmt.Println(split[len(split)-1])
			filtered_files = append(filtered_files, split[len(split)-1])
		}
	}

	return filtered_files
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
	// Set Variables u stands for use
	// Do in Main Controller
	u_bash :=  true
	u_dockerfile := true
	u_makefile := true
	
	// extension_filtered_files := 
	FilterFileExtensions(files, u_makefile, u_dockerfile, u_bash)
	// ShowSliceData(extension_filtered_files)
	
	// Test Connection
	ignore_file_parser.TestConnection()

	// Ignore Cases w Tokens
	// filter_cases := 
	ignore_file_parser.FilterFiles()
	
	// More
}
