package ignore_file_parser_test

import (
	"os"
	"testing"
	// "Soteria/ignore_file_parser"
)

func TestTestConnection(t *testing.T) {}

func TestShowSliceData(t *testing.T) {}

func TestFileToStringArray(t *testing.T) {}

func TestRemoveOneLineComments(t *testing.T) {}

func TestRemoveMultiLineComments(t *testing.T) {}

func TestRemoveAfterLineComments(t *testing.T) {}

func TestTokenize(t *testing.T) {}

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
