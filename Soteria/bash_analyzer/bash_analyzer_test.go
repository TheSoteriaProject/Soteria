package bash_analyzer_test

import (
	"os"
	"reflect"
	"testing"
	"path/filepath"
	"Soteria/bash_analyzer"
)

// Checks if it picks up deined variable defnitions. This can be checked based on all captiols.
func TestGetVariables(t *testing.T) {
	cwd, err := os.Getwd()
    	if err != nil {
        	t.Fatalf("Error getting current working directory: %v", err)
    	}

	// Test File
	test_file := "../../Files/sample_data/curl_examples.sh"
	
	// Construct the file path
    	file := filepath.Join(cwd, test_file)

	// Fucntion Call
	variable_list := bash_analyzer.GetVariables(file)

	// Test Data Variable define list
	test_list := []string{"DOWNLOAD_URL", 
			      "POST_URL", 
			      "POST_DATA", 
			      "param1", 
			      "param2",
			      "CURL",
			      "INSECURE", 
			      "DATA",
		      	      "HEADER", 
			      "REQUEST", 
			      "OUTPUT", 
			      "INSECURE_CURL", 
			      "command", 
			      "command"}


	// Compare
        if !reflect.DeepEqual(test_list, variable_list) {
        	t.Errorf("Expected: %v, Got: %v", test_list, variable_list)
    	}
}

func TestingController(t *testing.T) { 
        // TestConnections()
	TestGetVariables(t)
}       

func MainTest(m *testing.M) {
        os.Exit(m.Run())
} 
