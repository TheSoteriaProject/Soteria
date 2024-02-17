package diverter

import (
	"fmt"
	"path/filepath"
	"strings"
)

// TestConnection is used to Test Diverter Connection.
func TestConnection() {
	fmt.Println("Testing Diverter Connection.")
}

// DivertFiles is used to send ethe files to the correct static (analyzer || analyzers)
func DivertFiles(file_pool []string) {

	// Extract bash, Makefiles, and Dockerfiles to be used T or F
	// Set Variables u stands for use
	// Do in Main Controller
	u_bash := true
	u_dockerfile := true
	u_makefile := true

	for _, file := range file_pool {
		split := strings.Split(file, "/")
		extension := filepath.Ext(split[len(split)-1]) // MAY WANT FULL FILE PATH
		// first checks extension second checks filename and if equal to Makefile, Dockerfile, etc..
		// Makefile Check
		if u_makefile && strings.Contains(strings.ToLower(extension), strings.ToLower(".makefile")) || strings.Contains(strings.ToLower(split[len(split)-1]), strings.ToLower("makefile")) {
			fmt.Println("Diverted: " + file + " to Makefile Static Analyzer.")
		}
		// Dockerfile Check
		if u_dockerfile && strings.Contains(strings.ToLower(extension), strings.ToLower(".dockerfile")) || strings.Contains(strings.ToLower(split[len(split)-1]), strings.ToLower("dockerfile")) {
			fmt.Println("Diverted: " + file + " to Docker Static Analyzer.")
		}
		// Bash Check
		if u_bash && strings.Contains(strings.ToLower(extension), strings.ToLower(".sh")) {
			fmt.Println("Diverted: " + file + " to Bash Static Analyzer.")
		}
	}
}
