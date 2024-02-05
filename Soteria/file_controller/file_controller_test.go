// This file will be used to create test cases for expected outputs via JSOn or some other info in which we compare the info
// Let's consider test driven development NOT 100% certain this would be the best course of action.
// can't leave blank

package file_controller_test

import (
	// "fmt"
	"os"
	"io/ioutil"
	"strings"
	"testing"

	"Soteria/file_controller"
	//"Soteria/ignore_file_parser"
	//"Soteria/diverter"
	//"Soteria/lexers"
)


func TestConnections(t *testing.T) {
	// Redirect stdout to capture the output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Restore the original stdout after the test
	defer func() {
		os.Stdout = old
	}()

	file_controller.TestConnection()

	w.Close()
	capturedOutput, _ := ioutil.ReadAll(r)

	expect_from_file_controller := "Testing File Controller Connection."

	if strings.TrimSpace(string(capturedOutput)) != expect_from_file_controller {
		t.Errorf("Expected: %q, Got: %q", expect_from_file_controller, string(capturedOutput))
	}
}

func TestingController(t *testing.T) {
	// Run the tests
	TestConnections(t)
}

func TestMain(m *testing.M) {
	// Run the tests using the standard testing package
	os.Exit(m.Run())
}
