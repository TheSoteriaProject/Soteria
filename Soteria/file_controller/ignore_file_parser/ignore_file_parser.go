package ignore_file_parser

import (
	"fmt"
	"io/ioutil"
	"log"
	// "os"
)

func TestConnection() {
	fmt.Println("Testing Ignore File Parser Connection.")
}

func RemoveOneLineComments() {}

func RemoveMultiLineComments() {}

func Tokenize() []string { 
        return nil
}

func FilterFiles() []string {
	fmt.Println("Reading Ignore File")
	filename := "./.soteriaignore"
	content, err := ioutil.ReadFile(filename)
    	if err != nil {
        	log.Panicf("failed reading data from file: %s", err)
    	}

	fmt.Printf("Content: %s", content)

	// Remove Comments

	// Remove Multi-Line Comments

	// Tokenize

	return nil
}
