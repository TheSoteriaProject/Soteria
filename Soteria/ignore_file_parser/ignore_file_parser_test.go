package ignore_file_parser_test

import (
	"os"
	"testing"
	// "Soteria/ignore_file_parser"
)

func TestTestConnection(t *testing.T) {
	// t.Errorf("Expected: %q, Got: %q", "expected", "got")
}

func TestShowSliceData(t *testing.T) {
	// t.Errorf("Expected: %q, Got: %q", "expected", "got")
}

func TestFileToStringArray(t *testing.T) {
	// t.Errorf("Expected: %q, Got: %q", "expected", "got")
}

func TestRemoveOneLineComments(t *testing.T) {
	// t.Errorf("Expected: %q, Got: %q", "expected", "got")
}

func TestRemoveMultiLineComments(t *testing.T) {
	// t.Errorf("Expected: %q, Got: %q", "expected", "got")
}

func TestRemoveAfterLineComments(t *testing.T) {
	// t.Errorf("Expected: %q, Got: %q", "expected", "got")
}

func TestTokenize(t *testing.T) {
	// t.Errorf("Expected: %q, Got: %q", "expected", "got")
}

func TestFilterFiles(t *testing.T) {
	TestTestConnection(t)
	TestShowSliceData(t)
	TestFileToStringArray(t)
	TestRemoveOneLineComments(t)
	TestRemoveMultiLineComments(t)
	TestRemoveAfterLineComments(t)
	TestTokenize(t)
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
