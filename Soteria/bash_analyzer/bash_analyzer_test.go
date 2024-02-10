package bash_analyzer_test

import (
	"os"
	"reflect"
	"testing"
	"path/filepath"
	"Soteria/bash_analyzer"
)

// Checks if it picks up deined variable defnitions. This can be checked based on all captiols.
func TestGetVariables1(t *testing.T) {
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

// Checks if it picks up deined variable defnitions. This can be checked based on all captiols.
func TestGetVariables2(t *testing.T) {
        cwd, err := os.Getwd()
        if err != nil {
                t.Fatalf("Error getting current working directory: %v", err)
        }

        // Test File
        test_file := "../../Files/sample_data/ssh_examples.sh"
        
        // Construct the file path
        file := filepath.Join(cwd, test_file)

        // Fucntion Call
        variable_list := bash_analyzer.GetVariables(file)

        // Test Data Variable define list
        test_list := []string{"ssh_opts",
			      "StrictHostKeyChecking", 
			      "UserKnownHostsFile",
                              "ssh_cmd", 
                              "StrictHostKeyChecking", 
                              "UserKnownHostsFile"}

        // Compare
        if !reflect.DeepEqual(test_list, variable_list) {
                t.Errorf("Expected: %v, Got: %v", test_list, variable_list)
        }
}

// Checks if it picks up deined variable defnitions. This can be checked based on all captiols.
func TestGetVariables3(t *testing.T) {
        cwd, err := os.Getwd()
        if err != nil {
                t.Fatalf("Error getting current working directory: %v", err)
        }

        // Test File
        test_file := "../../Files/sample_data/wget_examples.sh"
        
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
                              "WGET",
                              "NO_CHECK_CERTIFICATE", 
                              "POST",
                              "HEADER", 
			      "header", 
                              "data", // sus????
                              "NO_CHECK_CERTIFICATE_WGET", 
                              "command", 
                              "command"}


        // Compare
        if !reflect.DeepEqual(test_list, variable_list) {
                t.Errorf("Expected: %v, Got: %v", test_list, variable_list)
        }
}

func TestGetVariableDefinitions1(t *testing.T) {
	cwd, err := os.Getwd()
        if err != nil {
                t.Fatalf("Error getting current working directory: %v", err)
        }

        // Test File
        test_file := "../../Files/sample_data/curl_examples.sh"
        
        // Construct the file path
        file := filepath.Join(cwd, test_file)

        // Fucntion Call
        definition_list := bash_analyzer.GetVariableDefinitions(file)

        // Test Data Definition define list
        test_list := []string{"'https://example.com/installer.pkg'", 
                              "'https://example.com/api/endpoint'", 
                              "'param1=value1&param2=value2'", 
                              "value1", 
                              "value2",
                              "'curl'",
                              "'--insecure'", 
                              "'--data'",
                              "'--header \"Content-Type: application/x-www-form-urlencoded\"'", 
                              "'--request'", 
                              "'--output'", 
                              "\"${CURL} ${INSECURE}\"", 
                              "$1", 
                              "('curl' '-k' '-o' 'installer3.pkg' \"${DOWNLOAD_URL}\")"}


        // Compare
        if !reflect.DeepEqual(test_list, definition_list) {
                t.Errorf("Expected: %v, Got: %v", test_list, definition_list)
        }
}

func TestingController(t *testing.T) { 
        // TestConnections()
	TestGetVariables1(t)
	TestGetVariableDefinitions1(t)
	TestGetVariables2(t)
        // TestGetVariableDefinitions2(t)
	TestGetVariables3(t)
        // TestGetVariableDefinitions3(t)
}

func MainTest(m *testing.M) {
        os.Exit(m.Run())
} 
