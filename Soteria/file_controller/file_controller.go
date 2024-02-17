package file_controller

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// TestConnection is used to Test File Controller Connection.
func TestConnection() {
	fmt.Println("Testing File Controller Connection.")
}

// ShowSliceData is to be used for Debugging Slices.
func ShowSliceData(path []string) {
	for _, path := range path {
		fmt.Printf("%s\n", path)
	}
}

// GetIgnoreDirs Gets Folders that should be Ignored.
func GetIgnoreDirs(path string) []string {
	ignoreDirs := []string{}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read each line and append it to the slice unless it is a comment
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(strings.TrimSpace(line), "#") {
			ignoreDirs = append(ignoreDirs, scanner.Text())
		}
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return nil
	}

	return ignoreDirs
}

// WalkTheFiles traverses the directory tree starting from the given project path and collects needed file paths.
func WalkTheFiles(path string, ignoreDirs []string) []string {
	files := make([]string, 0)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error walking the path:", err)
			return err
		}
		if info.IsDir() && ShouldSkipDir(path, ignoreDirs) {
			return filepath.SkipDir
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking the path:", err)
		return nil
	}
	return files
}

// ShouldSkipDir checks if the directory should be skipped based on the provided list.
func ShouldSkipDir(path string, ignoreDirs []string) bool {
	dir := filepath.Base(path)
	for _, ignoreDir := range ignoreDirs {
		if dir == ignoreDir {
			return true
		}
	}
	return false
}

// FilterFileExtensions checks each file to make sure it is the proper extension.
func FilterFileExtensions(files []string, u_makefile bool, u_dockerfile bool, u_bash bool) []string {
	filtered_files := []string{}

	for _, files := range files {
		split := strings.Split(files, "/")
		extension := filepath.Ext(split[len(split)-1]) // MAY WANT FULL FILE PATH
		// first checks extension second checks filename and if equal to Makefile, Dockerfile, etc..
		// Makefile Check
		if u_makefile && strings.Contains(strings.ToLower(extension), strings.ToLower(".makefile")) || strings.Contains(strings.ToLower(split[len(split)-1]), strings.ToLower("makefile")) {
			// fmt.Println(split[len(split)-1])
			filtered_files = append(filtered_files, files)
		}
		// Dockerfile Check
		if u_dockerfile && strings.Contains(strings.ToLower(extension), strings.ToLower(".dockerfile")) || strings.Contains(strings.ToLower(split[len(split)-1]), strings.ToLower("dockerfile")) {
			// fmt.Println(split[len(split)-1])
			filtered_files = append(filtered_files, files)
		}
		// Bash Check
		if u_bash && strings.Contains(strings.ToLower(extension), strings.ToLower(".sh")) {
			// fmt.Println(split[len(split)-1])
			filtered_files = append(filtered_files, files)
		}
	}

	return filtered_files
}

// FileController is the Main Controller and handles each step.
func FileController(path string) []string {
	ignore_file := "./.soteriaignore"
	ignoreDirs := GetIgnoreDirs(ignore_file)
	ShowSliceData(ignoreDirs)

	// Get All Files / Walk The Directories
	files := WalkTheFiles(path, ignoreDirs)

	// Extract bash, Makefiles, and Dockerfiles to be used T or F
	// Set Variables u stands for use
	// Do in Main Controller
	u_bash := true
	u_dockerfile := true
	u_makefile := true

	// Filter Down to just Bash, Makefiles, and Dockerfiles
	file_pool := FilterFileExtensions(files, u_makefile, u_dockerfile, u_bash)

	// Show Final File Pool to be Diverted
	ShowSliceData(file_pool)

	return file_pool
}
