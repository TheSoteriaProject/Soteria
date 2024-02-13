// Let's consider test driven development NOT 100% certain best option.

package file_controller_test

import (
	// "fmt"
	"io"
	"os"
	"strings"
	"testing"

	"Soteria/file_controller"
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
	capturedOutput, _ := io.ReadAll(r)

	expect_from_file_controller := "Testing File Controller Connection."

	if strings.TrimSpace(string(capturedOutput)) != expect_from_file_controller {
		t.Errorf("Expected: %q, Got: %q", expect_from_file_controller, string(capturedOutput))
	}
}

func TestShowSliceData(t *testing.T) {
	// Redirect stdout to capture the output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Restore the original stdout after the test
	defer func() {
		os.Stdout = old
	}()

	dummyStringSlice := []string{"apple", "banana", "cherry", "date", "elderberry"}

	file_controller.ShowSliceData(dummyStringSlice)

	w.Close()
	capturedOutput, _ := io.ReadAll(r)
	capturedOutputStr := strings.ReplaceAll(string(capturedOutput), "\r\n", "\n")

	expect_from_file_controller := "apple\nbanana\ncherry\ndate\nelderberry\n"
	expect_from_file_controller = strings.TrimSpace(expect_from_file_controller)
	if strings.TrimSpace(string(capturedOutputStr)) != expect_from_file_controller {
		t.Errorf("Expected: %q, Got: %q", expect_from_file_controller, string(capturedOutputStr))
	}
}

func TestWalkTheFiles(t *testing.T) {}

func TestGetAllFiles(t *testing.T) {
	// files, folders := file_controller.GetAllFilesAndFolders()
	// t.Errorf("Expected: %q, Got: %q", "expected", "got")
	// Needs a re-write
}

func TestFilterFileExtensions(t *testing.T) {
	// t.Errorf("Expected: %q, Got: %q", "expected", "got")
}

func TestCompareFiles(t *testing.T) {
	// t.Errorf("Expected: %q, Got: %q", "expected", "got")
}

func TestFileController(t *testing.T) {
	// Run the tests
	TestConnections(t)
	TestShowSliceData(t)
	TestWalkTheFiles(t)
	TestGetAllFiles(t)
	TestFilterFileExtensions(t)
	TestCompareFiles(t)
}

func TestMain(m *testing.M) {
	// Run the tests using the standard testing package
	os.Exit(m.Run())
}
