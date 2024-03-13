package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	// Custom Files
	"Soteria/diverter"
	"Soteria/file_controller"
	Help "Soteria/help"       // Alias Given
	JLogger "Soteria/logging" // Alias Given
)

func main() {
	fmt.Println("Welcome to Insecure Communication Linter.")

	// Confrim/Test File Controller Connection
	// file_controller.TestConnection()

	//Confirm/Test Diverter Connection
	// diverter.TestConnection()

	// if len(os.Args) > 1 {
	if flag.NFlag() >= 0 || flag.NArg() >= 0 {
		// Diasble Terminal Flags
		flag.Usage = func() {}

		// Test Flag
		var runTests bool
		flag.BoolVar(&runTests, "test", false, "Run tests")

		// Help Flag
		var helpUser bool
		flag.BoolVar(&helpUser, "help", false, "Help User")

		// Version Flag
		var versionCheck bool
		flag.BoolVar(&versionCheck, "version", false, "Version Check")
		flag.BoolVar(&versionCheck, "v", false, "Version Check")

		// Bash, Makefile, and Docker Flag
		// May be doing this wrong long term.
		enableBash := flag.Bool("enableBash", true, "Check Bash Files")
		enableMakefile := flag.Bool("enableMakefile", true, "Check Makefiles File")
		enableDockerfile := flag.Bool("enableDockerfile", true, "Check Dockerfile Files")

		// Logging Print Flag
		enableLogPrint := flag.Bool("enableLogPrint", true, "Check If Logs Print")

		// Warn Flag
		var warnUser bool
		flag.BoolVar(&warnUser, "warn", false, "Warn User")

		// Parse Flag
		flag.Parse()

		// Take Project Path
		args := flag.Args()
		input := ""
		if len(args) > 0 {
			input = args[len(args)-1]
		} else {
			input = flag.Arg(0)
		}

		if runTests {
			cmd := exec.Command("go", "test", "./...", "-v")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			fmt.Println("Running tests...")
			err := cmd.Run()
			if err != nil {
				fmt.Printf("Error running tests: %v\n", err)
				os.Exit(1)
			}
		} else if helpUser {
			Help.Info()
		} else if versionCheck {
			tagOutput := exec.Command("git", "rev-list", "--tags", "--max-count=1")
			outTag, errTag := tagOutput.Output()
			if errTag != nil {
				fmt.Println("Error: ", errTag)
			}

			tag := strings.TrimSpace(string(outTag))

			versionCmd := exec.Command("git", "describe", "--tags", tag)
			outVersion, errVersion := versionCmd.Output()
			if errVersion != nil {
				fmt.Println("Error:", errVersion)
				return
			}

			fmt.Println("Version:", strings.TrimSpace(string(outVersion)))
		} else {
			// Add Other flags
			// --warn etc...
			// Check for multiple flags possibly

			// If File Path Does Exist
			if _, err := os.Stat(input); err == nil {
				// All the Files that are to be checked.
				file_pool := file_controller.FileController(input, *enableMakefile, *enableDockerfile, *enableBash)
				// Divert Files to correct parser || parsers
				// Probably does not need the 'use' flags. Get Opinions.
				err := JLogger.DestroyJsonLog() // Truncates Old File
				if err != nil {
					fmt.Println("Error Destroying Logs.")
					os.Exit(1)
				}
				diverter.DivertFiles(file_pool, warnUser, *enableMakefile, *enableDockerfile, *enableBash, *enableLogPrint)
			} else if errors.Is(err, os.ErrNotExist) {
				// If Path Does Not Exist Throw Error and Exit
				// fmt.Println("Path Does Not Exist.")
				// Other Issue?
				fmt.Println("It seems you have given an invalid input. Try --help")
				os.Exit(1)
			} else {
				// If File Does Not Exist Throw Err and EXIT
				// SHOULD NOT MAKE IT HERE
				fmt.Println("File Does Not Exist")
				os.Exit(1)
			}
		}
	} else {
		// Invalid Input or Bad CLI Argument
		fmt.Println("It seems you have given an invalid input. Try --help")
	}
}
