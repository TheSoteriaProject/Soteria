package ignore_file_parser

import (
	"fmt"
	// "io"
	"bufio"
	// "log"
	"os"
	"strings"
)

func TestConnection() {
	fmt.Println("Testing Ignore File Parser Connection.")
}

func ShowSliceData(data []string) {
	line_number := 0
	for _, line := range data {
		fmt.Println(line_number, " : ", line)
		line_number += 1
	}
}

func FileToStringArray(filepath string) []string {
	readFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		if fileScanner.Text() != "" {
			fileLines = append(fileLines, fileScanner.Text())
		}
	}

	readFile.Close()

	return fileLines
}

func RemoveOneLineComments(filedata []string) []string {
	removedOneLiners := []string{}

	for _, line := range filedata {
		if !strings.HasPrefix(line, "#") {
			removedOneLiners = append(removedOneLiners, line)
		}
	}

	return removedOneLiners
}

func RemoveMultiLineComments(filedata []string) []string {
	removedMultiLiners := []string{}
	multiLineFlag := false

	for _, line := range filedata {
		if strings.HasPrefix(line, "|-start") {
			multiLineFlag = true
		}
		if strings.HasPrefix(line, "|-end") {
			multiLineFlag = false
		}

		if !multiLineFlag && !strings.HasPrefix(line, "|-end") {
			removedMultiLiners = append(removedMultiLiners, line)
		}
	}

	return removedMultiLiners
}

func RemoveAfterLineComments(filedata []string) []string {
	removedAfterLiners := []string{}
	commentIndex := 1

	for _, line := range filedata {
		for _, character := range line {
			if character != '#' {
				commentIndex += 1
			} else {
				break
			}
		}
		removedAfterLiners = append(removedAfterLiners, line[:commentIndex-1])
		commentIndex = 1
	}

	return removedAfterLiners
}

func Tokenize(filedata []string) []string {
	files_with_tokens := []string{}

	for _, line := range filedata {
		if strings.HasPrefix(strings.TrimSpace(line), "-") {
			if strings.HasSuffix(strings.TrimSpace(line), "/") {
				files_with_tokens = append(files_with_tokens, line+": IgnoreFolder")
			} else {
				files_with_tokens = append(files_with_tokens, line+": IgnoreFile")
			}
		} else if strings.HasPrefix(strings.TrimSpace(line), "+") {
			if strings.HasSuffix(strings.TrimSpace(line), "/") {
				files_with_tokens = append(files_with_tokens, line+": IncludeFolder")
			} else {
				files_with_tokens = append(files_with_tokens, line+": IncludeFile")
			}
		} else if strings.HasPrefix(strings.TrimSpace(line), "*") {
			files_with_tokens = append(files_with_tokens, line+": IgnoreExtension")
		} else {
			//
			fmt.Println("Error in Tokenize")
		}
	}

	return files_with_tokens
}

func FilterFiles() []string {
	// fmt.Println("Reading Ignore File.")
	filepath := "./.soteriaignore"

	// Read File into String Array
	fileLines := FileToStringArray(filepath)
	// ShowSliceData(fileLines) // Show Data

	// Remove Comments
	without_one_line_comments := RemoveOneLineComments(fileLines)
	// ShowSliceData(without_one_line_comments) // Show Data

	// Remove Multi-Line Comments
	without_multi_line_comments := RemoveMultiLineComments(without_one_line_comments)
	// ShowSliceData(without_multi_line_comments) // Show Data

	// Remove After Line Comments
	without_after_line_comments := RemoveAfterLineComments(without_multi_line_comments)
	// ShowSliceData(without_after_line_comments)

	// Tokenize
	tokenized_files := Tokenize(without_after_line_comments)
	// ShowSliceData(tokenized_files)

	return tokenized_files
}
