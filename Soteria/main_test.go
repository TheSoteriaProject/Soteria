package main_test

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

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

func TestMainController(t *testing.T) {
	TestMainNoInputs(t)
	TestMainHelpUserPage(t)
	TestMainVersionCheck(t)
}

func MainControllerTest(m *testing.M) {
	os.Exit(m.Run())
}
