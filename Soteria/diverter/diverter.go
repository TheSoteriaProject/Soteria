package diverter

import (
	"Soteria/bash_analyzer"
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// TestConnection is used to Test Diverter Connection.
func TestConnection() {
	fmt.Println("Testing Diverter Connection.")
}

// DivertFiles is used to send files to the correct static (analyzer || analyzers)
func DivertFiles(file_pool []string, warnUser bool, u_makefile bool, u_dockerfile bool, u_bash bool) {
	for _, file := range file_pool {
		split := strings.Split(file, "/")
		extension := filepath.Ext(split[len(split)-1]) // MAY WANT FULL FILE PATH
		// first checks extension second checks filename and if equal to Makefile, Dockerfile, etc..
		// Makefile Check
		if u_makefile && strings.Contains(strings.ToLower(extension), strings.ToLower(".makefile")) || strings.Contains(strings.ToLower(split[len(split)-1]), strings.ToLower("makefile")) {
			fmt.Println("Diverted: " + file + " to Makefile Static Analyzer.")
			// exec.Command("python", "example.py", file, strconv.FormatBool(warnUser)) // Change Name Later
		}
		// Dockerfile Check
		if u_dockerfile && strings.Contains(strings.ToLower(extension), strings.ToLower(".dockerfile")) || strings.Contains(strings.ToLower(split[len(split)-1]), strings.ToLower("dockerfile")) {
			fmt.Println("Diverted: " + file + " to Docker Static Analyzer.")
			exec.Command("python", "example.py", file, strconv.FormatBool(warnUser)) // Change Name Later
		}
		// Bash Check
		if u_bash && strings.Contains(strings.ToLower(extension), strings.ToLower(".sh")) {
			// fmt.Println("Diverted: " + file + " to Bash Static Analyzer.")
			// Can Pass via CLI However for this one it is written in go so I wont.
			bash_analyzer.BashController(file, warnUser)
		}
	}
}
