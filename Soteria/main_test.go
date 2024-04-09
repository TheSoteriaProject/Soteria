package main_test

import (
	"os"
	"testing"
)

/*
func TestMainNoInputs(t *testing.T) {
	cmd := exec.Command("./Soteria")
	cmd.Stderr = os.Stderr
	var cmdOut bytes.Buffer
	cmd.Stdout = &cmdOut

	err := cmd.Run()

	// Check if error
	if err != nil {
		// If the command returned an error, check if it's a non-zero exit status
		exitErr, ok := err.(*exec.ExitError)
		if !ok {
			t.Fatalf("Error running soteria: %v\n", err)
		}

		// If it's an ExitError, check the exit status
		if exitErr.ExitCode() != 1 {
			t.Fatalf("Error running soteria: %v\n", exitErr)
		}
	}

	// Expecting output from the command
	expectFromMainController := "Welcome to Insecure Communication Linter.\nIt seems you have given an invalid input. Try --help"

	// Compare expected and given
	if strings.TrimSpace(cmdOut.String()) != expectFromMainController {
		t.Errorf("Expected: %q, Got: %q", expectFromMainController, strings.TrimSpace(cmdOut.String()))
	}
}

func TestMainHelpUserPage(t *testing.T) {
	cmd := exec.Command("./Soteria", "--help")
	cmd.Stderr = os.Stderr
	var cmdOut bytes.Buffer
	cmd.Stdout = &cmdOut

	err := cmd.Run()

	// Check if error
	if err != nil {
		// If the command returned an error, check if it's a non-zero exit status
		exitErr, ok := err.(*exec.ExitError)
		if !ok {
			t.Fatalf("Error running soteria: %v\n", err)
		}

		// If it's an ExitError, check the exit status
		if exitErr.ExitCode() != 1 {
			t.Fatalf("Error running soteria: %v\n", exitErr)
		}
	}

	if cmdOut.String() == " " || cmdOut.String() == "" {
		t.Errorf("Printed Nothing.")
	}
}

func TestMainVersionCheck(t *testing.T) {
	cmd := exec.Command("./Soteria", "--version")
	cmd.Stderr = os.Stderr
	var cmdOut bytes.Buffer
	cmd.Stdout = &cmdOut

	err := cmd.Run()

	// Check if error
	if err != nil {
		// If the command returned an error, check if it's a non-zero exit status
		exitErr, ok := err.(*exec.ExitError)
		if !ok {
			t.Fatalf("Error running soteria: %v\n", err)
		}

		// If it's an ExitError, check the exit status
		if exitErr.ExitCode() != 1 {
			t.Fatalf("Error running soteria: %v\n", exitErr)
		}
	}

	if !strings.Contains(cmdOut.String(), "Welcome to Insecure Communication Linter.\nVersion: v") {
		t.Errorf("Version Not Found")
	}
}
*/

// Commented out not because they do not work but cause a github build action due to exec.
// So if wanting to text functionality uncomment this section and run it or modify possibly to create functions that can be called instead of what I did it.
func TestMainController(t *testing.T) {
	// TestMainNoInputs(t)
	// TestMainHelpUserPage(t)
	// TestMainVersionCheck(t)
}

func MainControllerTest(m *testing.M) {
	os.Exit(m.Run())
}
