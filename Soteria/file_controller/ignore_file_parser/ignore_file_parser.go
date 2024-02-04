package ignore_file_parser

import (
	"fmt"
	// "io/ioutil"
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

func fileToStringArray(filepath string) []string {
	readFile, err := os.Open(filepath)
        if err != nil {
                fmt.Println(err)
        }

	fileScanner := bufio.NewScanner(readFile)
    	fileScanner.Split(bufio.ScanLines)
    	var fileLines []string
  
    	for fileScanner.Scan() {
        	fileLines = append(fileLines, fileScanner.Text())
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
		
		if multiLineFlag == false && !strings.HasPrefix(line, "|-end") {
			removedMultiLiners = append(removedMultiLiners, line) 
		}
	}

	return removedMultiLiners
}

func Tokenize() []string { 
        return nil
}

func FilterFiles() []string {
	fmt.Println("Reading Ignore File.")
	filepath := "./.soteriaignore"
	
	// Read File into String Array
	fileLines := fileToStringArray(filepath)
	// ShowSliceData(fileLines) // Show Data

	// Remove Comments
	without_one_line_comments := RemoveOneLineComments(fileLines)
	// ShowSliceData(without_one_line_comments) // Show Data
	
	// Remove Multi-Line Comments
	without_multi_line_comments := RemoveMultiLineComments(without_one_line_comments)
	ShowSliceData(without_multi_line_comments)
	
	// Tokenize

	return nil
}
