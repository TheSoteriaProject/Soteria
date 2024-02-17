package diverter

import (
	"fmt"
)

// TestConnection is used to Test Diverter Connection.
func TestConnection() {
	fmt.Println("Testing Diverter Connection.")
}

// DivertFiles is used to send ethe files to the correct static (analyzer || analyzers)
func DivertFiles(file_pool []string) ([]string, error) {
	// Do Nothing
	return file_pool, nil
}
