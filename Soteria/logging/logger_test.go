package logger_test

import (
	"os"
	"testing"
)

func TestLoggerController(t *testing.T) {
	// Not Needed Currently.
	// Added so if needed in the future it is here.
	// Should Probably be used, but based on time determines if this will be added.
}

func LoggerControllerTest(m *testing.M) {
	os.Exit(m.Run())
}
