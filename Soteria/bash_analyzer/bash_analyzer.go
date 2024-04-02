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

func ReadLines(_file string, filename string, warnUser bool, section string, command string, enableLogPrint bool) {
	file, err := os.Open(_file)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	lineTemp := ""
	lineTempStartCounter := -1

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(strings.ToLower(line), "#") && !strings.HasPrefix(strings.ToLower(line), "echo") {
			pattern1 := "\\b" + section + "\\b"
			pattern2 := command + "\\b"
			re1 := regexp.MustCompile(pattern1)
			re2 := regexp.MustCompile(pattern2)

			// Check if the text contains the exact matches for both patterns
			if re1.MatchString(line) && re2.MatchString(line) {
				// Adjust this to be less stupid
				ErrorType := "Error"
				if warnUser {
					ErrorType = "Warn"
				}

				// Deal with None Error cases like comments and echos
				// Doeals with other cases that would create a bad temp file. So basically this is how I dealt with not screwing up the temp file check.
				// Allowing me to handle these extra cases without breaking the others.
				if !strings.HasPrefix(strings.ToLower(line), "#") && !strings.HasPrefix(strings.ToLower(line), "echo") {
					if strings.Contains(line, "Ignore Match") { // Ignore Line Case
						ErrorType = "Warn"
					}

					// Deal with ; and \ but issue need to check for ; in the second
					if strings.HasSuffix(line, "\\") || strings.HasSuffix(lineTemp, "\\") {
						lineTempStartCounter += 1
						// Need to also add ; check.
						// Should I Remove the newline on 'fake' compile???
						lineTemp += line
						if !strings.HasSuffix(line, "\\") {
							// I Dont like how it logs this, but any more complex logic may break something before the EOS.
							// I personally think create a link count logic and use it to hold start and finish and show the range of the issues.

							// This handles ; but could probably be better but it works.
							if strings.Contains(lineTemp, ";") {
								lines := strings.Split(lineTemp, ";")
								for _, line := range lines {
									// Line Number is off unless handled something like this
									JLogger.JsonLogger(filename, lineNumber-lineTempStartCounter, line, section+" "+command, ErrorType, enableLogPrint)
								}
							} else {
								JLogger.JsonLogger(filename, lineNumber, lineTemp, section+" "+command, ErrorType, enableLogPrint)
							}
							lineTemp = ""
							lineTempStartCounter = -1
						}
					} else if strings.Contains(line, ";") {
						lines := strings.Split(line, ";")
						for _, line := range lines {
							JLogger.JsonLogger(filename, lineNumber, line, section+" "+command, ErrorType, enableLogPrint)
						}
					} else {
						JLogger.JsonLogger(filename, lineNumber, line, section+" "+command, ErrorType, enableLogPrint)
					}
				}
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
		regex_pattern := `\b([a-zA-Z_][a-zA-Z0-9_]*)\s*=['"($]` //`\b([A-Z_][A-Z0-9_]*)\s*=\s*`
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
		// regex_pattern := `(=)([('].*)` // Seems Better
		regex_pattern := `(=)([('].*)` // `[A-Z]\b(=)([('].*)`
		regex_expression := regexp.MustCompile(regex_pattern)
		matches := regex_expression.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if len(match) > 1 {
				match2 := strings.Replace(match[2], " # Ignore Match", "", -1) // Dumb AF need to adjust. Spacing Everything garbage. Just for testing I guess.
				definition_list = append(definition_list, match2)
			}
		}

		// Second Pass
		// regex_pattern = `\b=([a-zA-Z_\/:-][a-zA-Z0-9_\/:-]*)`
		regex_pattern = `[a-zA-Z](=)(["${].*|[$[a-z]].*)` // `((?:^|[\-])*)([a-zA-Z_][a-zA-Z0-9_]*)=["${]` // `(?<!-)([a-zA-Z_][a-zA-Z0-9_]*)=["${]` // `[A-Z]\b(=)("\$\{[^"]+?\}")`
		regex_expression = regexp.MustCompile(regex_pattern)
		matches2 := regex_expression.FindAllStringSubmatch(line, -1)

		for _, match2 := range matches2 {
			if len(match2) > 1 {
				definition_list = append(definition_list, match2[2])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error Reaading: ", err)
	}

	return definition_list
}

// SwapLine takes the line with the variable possibilities and checks if defined.
func SwapLine(line string, variables []string, definitions []string) string {
	// fmt.Println("Line: ", line) // debug
	if len(definitions) == len(variables) {
		for i, variable := range variables {
			if strings.Contains(line, "\"${"+variable+"}\"") {
				line = strings.Replace(line, "\"${"+variable+"}\"", definitions[i], -1)
				line = SwapLine(line, variables, definitions)
			} else if strings.Contains(line, "${"+variable+"}") {
				line = strings.Replace(line, "${"+variable+"}", definitions[i], -1)
				line = SwapLine(line, variables, definitions)
			} else if strings.Contains(line, "$"+variable+"") {
				line = strings.Replace(line, "$"+variable+"", definitions[i], -1)
				line = SwapLine(line, variables, definitions)
			} else {
				continue
			}
			// fmt.Println(variable, " : ", line, " : ", definitions[i])
			// add other case???
		}
	}

	return line
}

// VariableSwap swaps the variables with what they were defined with in the code.
func VariableSwap(file string, warnUser bool, variables []string, variable_definitions []string) {
	// fmt.Println("File: ", file) // Debug
	oldFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer oldFile.Close()

	newFile, err := os.Create("../Soteria/bash_analyzer/temp.sh")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer newFile.Close()

	scanner := bufio.NewScanner(oldFile)
	writer := bufio.NewWriter(newFile)

	for scanner.Scan() {
		line := scanner.Text()
		// if !strings.HasPrefix(strings.ToLower(line), "#") && !strings.HasPrefix(strings.ToLower(line), "echo") {
		swappedLine := SwapLine(line, variables, variable_definitions) + "\n"
		_, err := writer.WriteString(swappedLine)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// fmt.Println(line)
		// }
	}
	// Flush
	writer.Flush()
}

// CheckForHiddenInsecureCommunication from the file.
func CheckForHiddenInsecureCommunication(filepath string, warn_file string, warnUser bool, variables []string, variable_definitions []string, enableLogPrint bool) {
	filename := filepath

	/* Not Needed. Leave for now.
	yamlData, err := ReadYAMLFile(warn_file)
	if err != nil {
		log.Fatalf("error reading YAML file: %v", err)
	}

	var data map[string]interface{}

	if err := yaml.Unmarshal([]byte(yamlData), &data); err != nil {
		log.Fatalf("error: %v", err)
	}
	*/
	// fmt.Println("Inside CheckForHiddenInsecureCommunication")
	VariableSwap(filepath, warnUser, variables, variable_definitions)

	// Deal with hardcoded filename
	CheckForInsecureCommunication("../Soteria/bash_analyzer/temp.sh", filename, warnUser, warn_file, variables, variable_definitions, enableLogPrint)
}

func CheckForInsecureCommunication(filepath string, filename string, warnUser bool, warn_file string, variables []string, variable_definitions []string, enableLogPrint bool) {
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
			ReadLines(filepath, filename, warnUser, section, command.(string), enableLogPrint)
		}
	}
}

func BashController(file string, warnUser bool, enableLogPrint bool) {
	// Pass File Name/Path
	v := GetVariables(file)
	vd := GetVariableDefinitions(file)
	// ShowTwoSlicesData(v, vd)

	// Iterate YAML
	warn_file := "bash_analyzer/warn.yaml"

	// CheckForInsecureCommunication(file, warnUser, warn_file, v, vd) // V and D probably useless for in-line

	CheckForHiddenInsecureCommunication(file, warn_file, warnUser, v, vd, enableLogPrint)
}
