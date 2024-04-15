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

	// Open File for Ignore Directories
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read each line and append it to the slice unless it is "Commnted out"
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(strings.TrimSpace(line), "#") {
			ignoreDirs = append(ignoreDirs, scanner.Text())
		}
	}

	// Check if errors occured
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return nil
	}

	return ignoreDirs
}

// WalkTheFiles traverses the directory tree starting from the given project path and collects needed file paths.
func WalkTheFiles(path string, ignoreDirs []string) []string {
	files := []string{}

	// Use pre-made function from path/filepath that walks the given roject path.
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error walking the path:", err)
			return err
		}
		// If it is a directory and depending on Skip Directory Status depends if called again.
		if info.IsDir() && ShouldSkipDir(path, ignoreDirs) {
			return filepath.SkipDir
		}
		// If file append
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	// If error walking path
	if err != nil {
		fmt.Println("Error walking the path:", err)
		return nil
	}
	return files
}

// ShouldSkipDir checks if the directory should be skipped based on the provided list.
func ShouldSkipDir(path string, ignoreDirs []string) bool {
	dir := filepath.Base(path)

	// Check if found in the ignore directory slice
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

	// Loop File Pool
	for _, file := range files {
		split := strings.Split(file, "/")
		extension := filepath.Ext(split[len(split)-1])

		// Makefile Check (Every part checks if enabled. The second part is checking the . extension or naming.)
		if (enableMakefile && strings.Contains(strings.ToLower(extension), strings.ToLower(".makefile"))) || (enableMakefile && strings.Contains(strings.ToLower(split[len(split)-1]), strings.ToLower("makefile"))) {
			filtered_files = append(filtered_files, file)
		}
		// Dockerfile Check (Every part checks if enabled. The second part is checking the . extension or naming.)
		if (enableDockerfile && strings.Contains(strings.ToLower(extension), strings.ToLower(".dockerfile"))) || (enableDockerfile && strings.Contains(strings.ToLower(split[len(split)-1]), strings.ToLower("dockerfile"))) {
			filtered_files = append(filtered_files, file)
		}
		// Bash Check  (Every part checks if enabled. The second part is checking the . extension or naming and also for this one checks the first two bytes for the #! in hex for probably no reason.)
		if (enableBash && strings.Contains(strings.ToLower(extension), strings.ToLower(".sh"))) || (enableBash && strings.Contains(strings.ToLower(split[len(split)-1]), strings.ToLower("-sh "))) || (enableBash && isFileBash(file)) {
			filtered_files = append(filtered_files, file)
		}
	}

	return filtered_files
}

// FileController is the Main Controller and handles each step.
func FileController(path string, enableMakefile bool, enableDockerfile bool, enableBash bool) []string {
	ignore_file := "Soteria/.soteriaignore" // Path to ignore file that contains the block directories.
	ignoreDirs := GetIgnoreDirs(ignore_file)

	// Get All Files / Walk The Directories
	files := WalkTheFiles(path, ignoreDirs)

	// Filter Down to just Bash, Makefiles, and Dockerfiles
	file_pool := FilterFileExtensions(files, enableMakefile, enableDockerfile, enableBash)

	return file_pool
}
