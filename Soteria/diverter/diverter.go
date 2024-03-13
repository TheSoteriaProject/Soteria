package diverter

import (
	"Soteria/bash_analyzer"
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// TestConnection is used to Test Diverter Connection.
func TestConnection() {
	fmt.Println("Testing Diverter Connection.")
}

// Maybe Make a file that is access as a helper for all if needed.
// ReadYAMLFile reads YAML data from a file and returns it as a byte slice.
func ReadYAMLFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// fileContainsInsecureCommunication is used to check if the file is insecure possibly.
func FileContainsInsecureCommunication(file string, warnFile string) bool {
	File, err := os.Open(file)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	defer File.Close()

	scanner := bufio.NewScanner(File)

	yamlData, err := ReadYAMLFile(warnFile)
	if err != nil {
		log.Fatalf("error reading YAML file: %v", err)
	}

	var data map[string]interface{}

	if err := yaml.Unmarshal([]byte(yamlData), &data); err != nil {
		log.Fatalf("error: %v", err)
	}

	for scanner.Scan() {
		line := scanner.Text()
		for section, commands := range data {
			for _, command := range commands.([]interface{}) {
				// Should be and?? Test more later.
				if strings.Contains(line, command.(string)) && strings.Contains(line, section) {
					return true
				}
			}
		}
	}
	return false
}

// DivertFiles is an extension of DivertFile which is done as a pre-check to see if the file should be scanned more in depth.
func DivertFiles(file_pool []string, warnUser bool, u_makefile bool, u_dockerfile bool, u_bash bool, enableLogPrint bool) {
	warn_file := "/Users/logangarrett03/Desktop/git/Soteria/Soteria/bash_analyzer/warn.yaml" // Bad Path Fix before merge.
	for _, file := range file_pool {

		if !FileContainsInsecureCommunication(file, warn_file) {
			continue
		} else {

			done := make(chan bool)
			// Start the function in a goroutine
			go func() {
				DivertFile(file, warnUser, u_makefile, u_dockerfile, u_bash, enableLogPrint)
				// Signal completion by sending true to the channel
				done <- true
			}()

			// Wait for either the function to complete or the timeout
			select {
			case <-done:
				// fmt.Println("Function completed before timeout")
				continue
			case <-time.After(60 * time.Second):
				fmt.Println("Function timed out")
			}
		}
	}
}

// DivertFiles is used to send files to the correct static (analyzer || analyzers)
func DivertFile(file string, warnUser bool, u_makefile bool, u_dockerfile bool, u_bash bool, enableLogPrint bool) {
	// for _, file := range file_pool {
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
	if u_bash && strings.Contains(strings.ToLower(extension), strings.ToLower(".sh")) || u_bash && strings.Contains(strings.ToLower(split[len(split)-1]), "-sh") {
		fmt.Println("Diverted: " + file + " to Bash Static Analyzer.")
		// Can Pass via CLI However for this one it is written in go so I wont.
		bash_analyzer.BashController(file, warnUser, enableLogPrint)
	}
	// }
}

/*
// DivertFiles is used to send files to the correct static (analyzer || analyzers)
func DivertFiles(file_pool []string, warnUser bool, u_makefile bool, u_dockerfile bool, u_bash bool, enableLogPrint bool) {
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
		if (u_bash && strings.Contains(strings.ToLower(extension), strings.ToLower(".sh"))) || (u_bash && strings.Contains(strings.ToLower(split[len(split)-1]), strings.ToLower("-sh "))) {
			fmt.Println("Diverted: " + file + " to Bash Static Analyzer.")
			// Can Pass via CLI However for this one it is written in go so I wont.
			bash_analyzer.BashController(file, warnUser, enableLogPrint)
		}
	}
}
*/
