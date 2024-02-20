package bash_analyzer

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	JLogger "Soteria/logging"

	"gopkg.in/yaml.v3"
)

// ShowSliceData is to be used for Debugging Slices.
func ShowSliceData(path []string) {
	for _, path := range path {
		fmt.Printf("%s\n", path)
	}
}

// ShowTwoSlicesData is used to iterate over two slices
func ShowTwoSlicesData(slice1 []string, slice2 []string) {
	minLength := len(slice1)
	if len(slice2) < minLength {
		minLength = len(slice2)
	}

	for i := 0; i < minLength; i++ {
		fmt.Println(slice1[i] + " : " + slice2[i])
	}
}

// ReadYAMLFile reads YAML data from a file and returns it as a byte slice.
func ReadYAMLFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ReadLines(filename string, warnUser bool, section string, command string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(strings.ToLower(line), "#") && !strings.HasPrefix(strings.ToLower(line), "echo") {
			pattern1 := "\\b" + section + "\\b"
			pattern2 := command
			re1 := regexp.MustCompile(pattern1)
			re2 := regexp.MustCompile(pattern2)

			// Check if the text contains the exact matches for both patterns
			if re1.MatchString(line) && re2.MatchString(line) {
				// fmt.Println("Found: ", line+"\nContains: "+section+" "+command)
				ErrorType := "Error"
				if warnUser {
					ErrorType = "Warn"
				}
				JLogger.JsonLogger(filename, lineNumber, line, section+" "+command, ErrorType)
			}
		}
		lineNumber += 1
	}

	if err := scanner.Err(); err != nil {
		return
	}
}

// GetVariables reads the variables from the Bash File.
func GetVariables(file_name string) []string {
	variable_list := []string{}

	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println("Error Opening: ", err)
		return nil
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// check regex
		// grab line number ???
		regex_pattern := `\b([a-zA-Z_][a-zA-Z0-9_]*)\s*=\s*`
		regex_expression := regexp.MustCompile(regex_pattern)
		matches := regex_expression.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if len(match) > 1 {
				variable_list = append(variable_list, match[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error Reaading: ", err)
	}

	return variable_list
}

// GetVariableDefnitions reads the variable definitions from the Bash File.
func GetVariableDefinitions(file_name string) []string {
	definition_list := []string{}
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println("Error Opening: ", err)
		return nil
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// check regex
		// grab line number ???
		regex_pattern := `(=)([($'"].*)` // Seems Better
		regex_expression := regexp.MustCompile(regex_pattern)
		matches := regex_expression.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if len(match) > 1 {
				definition_list = append(definition_list, match[2])
			}
		}

		// Second Pass
		regex_pattern = `\b=([a-zA-Z_\/:-][a-zA-Z0-9_\/:-]*)`
		regex_expression = regexp.MustCompile(regex_pattern)
		matches2 := regex_expression.FindAllStringSubmatch(line, -1)

		for _, match2 := range matches2 {
			if len(match2) > 1 {
				definition_list = append(definition_list, match2[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error Reaading: ", err)
	}

	return definition_list
}

// SwapLine takes the line with the variable possibilities and checks if defined.
func SwapLine(line string, variables []string, variable_definitions []string) string {
	newline := ""
	minLength := len(variables)
	if len(variable_definitions) < minLength {
		minLength = len(variable_definitions)
	}

	if len(strings.TrimSpace(line)) > 0 {
		for i := 0; i < minLength; i++ {
			fmt.Println("-----------------------------------------------------")
			fmt.Println(line)
			fmt.Println(variables[i] + " : " + variable_definitions[i])
			fmt.Println("-----------------------------------------------------")
		}
	}

	return newline
}

// VariableSwap swaps the variables with what they were defined with in the code.
func VariableSwap(file string, variables []string, variable_definitions []string) {
	oldFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer oldFile.Close()

	newFile, err := os.Create("temp.sh")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer newFile.Close()

	scanner := bufio.NewScanner(oldFile)
	writer := bufio.NewWriter(newFile)

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(strings.ToLower(line), "#") && !strings.HasPrefix(strings.ToLower(line), "echo") {
			swappedLine := SwapLine(scanner.Text(), variables, variable_definitions)
			_, err := writer.WriteString(swappedLine)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			// fmt.Println(line)
		}
	}
	// Flush
	writer.Flush()
}

/*
// CheckForHiddenInsecureCommunication from the file.
func CheckForHiddenInsecureCommunication(filepath string, variables []string, variable_definitions []string) {
	yamlData, err := ReadYAMLFile(filepath)
	if err != nil {
		log.Fatalf("error reading YAML file: %v", err)
	}

	var data map[string]interface{}

	if err := yaml.Unmarshal([]byte(yamlData), &data); err != nil {
		log.Fatalf("error: %v", err)
	}

	// Super Bad Code
	for section, commands := range data {
		// fmt.Println(section)
		for _, variable_definition := range variable_definitions {
			// -1 to drop :
			fmt.Println(section, variable_definition)
			if section == variable_definition {
				for _, command := range commands.([]interface{}) {
					fmt.Println("  ", command)
				}
			}
		}
	}
} */

func CheckForInsecureCommunication(filepath string, warnUser bool, warn_file string, variables []string, variable_definitions []string) {
	yamlData, err := ReadYAMLFile(warn_file)
	if err != nil {
		log.Fatalf("error reading YAML file: %v", err)
	}

	var data map[string]interface{}

	if err := yaml.Unmarshal([]byte(yamlData), &data); err != nil {
		log.Fatalf("error: %v", err)
	}

	for section, commands := range data {
		for _, command := range commands.([]interface{}) {
			// fmt.Println(section, command)
			// Just grep file if not echo or comment??? INLINE ONLY
			ReadLines(filepath, warnUser, section, command.(string))
		}
	}
}

func BashController(file string, warnUser bool) {
	// Pass File Name/Path
	v := GetVariables(file)
	vd := GetVariableDefinitions(file)
	// ShowTwoSlicesData(v, vd)

	// Iterate YAML
	warn_file := "bash_analyzer/warn.yaml"

	// CheckForHiddenInsecureCommunication(warn_file, v, vd)
	// VariableSwap(file, v, vd)

	CheckForInsecureCommunication(file, warnUser, warn_file, v, vd) // V and D probably useless for in-line
}
