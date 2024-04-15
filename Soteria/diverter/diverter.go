package diverter

import (
	"Soteria/bash_analyzer"
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
func DivertFiles(file_pool []string, warnUser bool, enableMakefile bool, enableDockerfile bool, enableBash bool, enableLogPrint bool) {
	warn_file := "../Soteria/bash_analyzer/rules.yaml" // Thought everyone was gonne use same format and remove later, but I guess not....
	for _, file := range file_pool {

		if !FileContainsInsecureCommunication(file, warn_file) {
			continue
		} else {

			done := make(chan bool)
			// Start the function in a goroutine
			go func() {
				DivertFile(file, warnUser, enableMakefile, enableDockerfile, enableBash, enableLogPrint)
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
func DivertFile(file string, warnUser bool, enableMakefile bool, enableDockerfile bool, enableBash bool, enableLogPrint bool) {
	split := strings.Split(file, "/")
	extension := filepath.Ext(split[len(split)-1]) // MAY WANT FULL FILE PATH
	// first checks extension second checks filename and if equal to Makefile, Dockerfile, etc..
	// Makefile Check
	if enableMakefile && strings.Contains(strings.ToLower(extension), strings.ToLower(".makefile")) || strings.Contains(strings.ToLower(split[len(split)-1]), strings.ToLower("makefile")) {
		fmt.Println("Diverted: " + file + " to Makefile Static Analyzer.")
		cmd := exec.Command("python3", "makeFileLinter/makefilelinter.py", file) // Change Name Later

		// Run the command and capture its output and any potential errors
		output, _ := cmd.CombinedOutput()

		// Print the output
		if enableLogPrint {
			fmt.Println(string(output))
		}
	}
	// Dockerfile Check
	if enableDockerfile && strings.Contains(strings.ToLower(extension), strings.ToLower(".dockerfile")) || strings.Contains(strings.ToLower(split[len(split)-1]), strings.ToLower("dockerfile")) {
		fmt.Println("Diverted: " + file + " to Docker Static Analyzer.")
		fmt.Println("Diverted: " + file + " to Makefile Static Analyzer.")
		cmd := exec.Command("python3", "dockerfile_linter/dockerfile_linter.py", file) // Change Name Later

		// Run the command and capture its output and any potential errors
		output, _ := cmd.CombinedOutput()

		// Print the output
		if enableLogPrint {
			fmt.Println(string(output))
		}
	}
	// Bash Check
	if enableBash && strings.Contains(strings.ToLower(extension), strings.ToLower(".sh")) || enableBash && strings.Contains(strings.ToLower(split[len(split)-1]), "-sh") {
		fmt.Println("Diverted: " + file + " to Bash Static Analyzer.")
		// Can Pass via CLI However for this one it is written in go so I wont.
		bash_analyzer.BashController(file, warnUser, enableLogPrint)
	}
}
