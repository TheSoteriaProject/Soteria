package bash_analyzer

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	//"Soteria/logging"
)

// ShowSliceData is to be used for Debugging Slices.
func ShowSliceData(path []string) {
	for _, path := range path {
		fmt.Printf("%s\n", path)
	}
}

// ShowTwoSlicesData is used to iterate over two slices
func ShowTwoSlicesData(slice1 []string, slice2 []string) {
	// Ensure both slices have the same length
	minLength := 0
	if len(slice2) == len(slice1) {
		minLength = len(slice2)
	} else {
		fmt.Println("Error. Length of slices must be equal.")
		// Find a way to deal with this.
	}

	// Iterate over slices and print side by side
	for i := 0; i < minLength; i++ {
		fmt.Println(slice1[i] + " : " + slice2[i])
	}
}

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

func BashController(file string) {
	// Pass File Name/Path
	temp := GetVariables(file)
	temp2 := GetVariableDefinitions(file)
	ShowTwoSlicesData(temp, temp2)
}
