package bash_analyzer_test

import (
	"os"
	"reflect"
	"testing"
	"Soteria/bash_analyzer"
)

// Checks if it picks up deined variable defnitions. This can be checked based on all captiols.
func TestGetVariables(t *testing.T) {
	// Test File
	file := "/Soteria/Files/sample_data/curl_examples.sh"
	
	// Fucntion Call
	variable_list := bash_analyzer.GetVariables(file)

	// Test Data
	// Gonna have to deal with POST since it is not a variable but is used so reserved list??
	test_list := []string{"DOWNLOAD_URL", 
			      "POST_URL", 
			      "POST_DATA",
			      "CURL",
			      "INSECURE", 
			      "DATA",
		      	      "HEADER", 
			      "REQUEST", 
			      "OUTPUT", 
			      "INSECURE_CURL"}

	// local variables will be a pain

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
