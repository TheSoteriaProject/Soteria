package bash_analyzer

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	//"Soteria/logging"
)

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

func BashController() {
	// Dummy Pass in for now
	GetVariables("./nothing.txt")
}
