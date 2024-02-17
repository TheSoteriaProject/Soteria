// Let's consider test driven development NOT 100% certain best option.

package file_controller_test

import (
	// "fmt"
	"io"
	"os"
	"reflect"
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

/*
func TestWalkTheFiles(t *testing.T) {
	path := "../../Files"
	ignore_Dirs := []string{}
	files := file_controller.WalkTheFiles(path, ignore_Dirs)

	test_list := []string{
		"../../Files/DoNotEnterFile.txt",
		"../../Files/DoNotEnterFolder/dummy.txt",
		"../../Files/Dockerfile",
		"../../Files/Makefile",
		"../../Files/NonBadFileCOBOL.cob",
		"../../Files/NonBadFilePEARL.pl",
		"../../Files/badBash.sh",
		"../../Files/sample_data/curl_examples.Dockerfile",
		"../../Files/sample_data/curl_examples.Makefile",
		"../../Files/sample_data/curl_examples.sh",
		"../../Files/sample_data/ssh_examples.Makefile",
		"../../Files/sample_data/ssh_examples.sh",
		"../../Files/sample_data/wget_examples.Dockerfile",
		"../../Files/sample_data/wget_examples.Makefile",
		"../../Files/sample_data/wget_examples.sh"}

	// Compare
	if !reflect.DeepEqual(test_list, files) {
		t.Errorf("Expected: %v, Got: %v", test_list, files)
	}
} */

func TestFilterFileExtensions(t *testing.T) {
	// Datset
	files_list := []string{
		"../../Files/sample_data/curl_examples.pearl",
		"../../Files/sample_data/curl_examples.txt",
		"../../Files/sample_data/curl_examples.Dockerfile",
		"../../Files/sample_data/curl_examples.Makefile",
		"../../Files/sample_data/curl_examples.sh",
		"../../Files/sample_data/ssh_examples.Makefile",
		"../../Files/sample_data/ssh_examples.sh",
		"../../Files/sample_data/curl_examples.COBOL",
		"../../Files/sample_data/wget_examples.Dockerfile",
		"../../Files/sample_data/wget_examples.Makefile",
		"../../Files/sample_data/wget_examples.sh"}

	// call function
	files := file_controller.FilterFileExtensions(files_list, true, true, true)

	// Test Data
	test_list := []string{
		"../../Files/sample_data/curl_examples.Dockerfile",
		"../../Files/sample_data/curl_examples.Makefile",
		"../../Files/sample_data/curl_examples.sh",
		"../../Files/sample_data/ssh_examples.Makefile",
		"../../Files/sample_data/ssh_examples.sh",
		"../../Files/sample_data/wget_examples.Dockerfile",
		"../../Files/sample_data/wget_examples.Makefile",
		"../../Files/sample_data/wget_examples.sh"}

	// Compare
	if !reflect.DeepEqual(test_list, files) {
		t.Errorf("Expected: %v, Got: %v", test_list, files)
	}
}

func TestFileController(t *testing.T) {
	// Run the tests
	TestConnections(t)
	TestShowSliceData(t)
	// TestWalkTheFiles(t)
	// TestFilterFileExtensions(t)
}

func TestMain(m *testing.M) {
	// Run the tests using the standard testing package
	os.Exit(m.Run())
}
