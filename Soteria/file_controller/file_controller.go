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

// isFileBash Checks if the file contains the shebang if it doesnt end with .sh or have -sh
func isFileBash(filepath string) bool {
	// Open File from Filepath
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return false
	}
	defer file.Close()

	// Grab First Two Bytes
	buf := make([]byte, 2)
	n, err := file.Read(buf)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return false
	}

	// Check for shebang
	if n >= 2 && buf[0] == '#' && buf[1] == '!' {
		return true
	} else {
		return false
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
func FilterFileExtensions(files []string, enableMakefile bool, enableDockerfile bool, enableBash bool) []string {
	filtered_files := []string{}

	for _, file := range files {
		split := strings.Split(file, "/")
		extension := filepath.Ext(split[len(split)-1]) // MAY WANT FULL FILE PATH
		// first checks extension second checks filename and if equal to Makefile, Dockerfile, etc..
		// Makefile Check
		if (enableMakefile && strings.Contains(strings.ToLower(extension), strings.ToLower(".makefile"))) || (enableMakefile && strings.Contains(strings.ToLower(split[len(split)-1]), strings.ToLower("makefile"))) {
			// fmt.Println(split[len(split)-1])
			filtered_files = append(filtered_files, file)
		}
		// Dockerfile Check
		if (enableDockerfile && strings.Contains(strings.ToLower(extension), strings.ToLower(".dockerfile"))) || (enableDockerfile && strings.Contains(strings.ToLower(split[len(split)-1]), strings.ToLower("dockerfile"))) {
			// fmt.Println(split[len(split)-1])
			filtered_files = append(filtered_files, file)
		}
		// Bash Check
		if (enableBash && strings.Contains(strings.ToLower(extension), strings.ToLower(".sh"))) || (enableBash && strings.Contains(strings.ToLower(split[len(split)-1]), strings.ToLower("-sh "))) || (enableBash && isFileBash(file)) {
			// fmt.Println(split[len(split)-1])
			filtered_files = append(filtered_files, file)
		}
	}

	return filtered_files
}

// FileController is the Main Controller and handles each step.
func FileController(path string, enableMakefile bool, enableDockerfile bool, enableBash bool) []string {
	ignore_file := "./.soteriaignore"
	ignoreDirs := GetIgnoreDirs(ignore_file)
	// ShowSliceData(ignoreDirs)

	// Get All Files / Walk The Directories
	files := WalkTheFiles(path, ignoreDirs)

	// Filter Down to just Bash, Makefiles, and Dockerfiles
	file_pool := FilterFileExtensions(files, enableMakefile, enableDockerfile, enableBash)

	// Show Final File Pool to be Diverted
	// ShowSliceData(file_pool)

	return file_pool
}
