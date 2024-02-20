package help

import (
	"fmt"
)

func Info() {
	fmt.Printf("\n") // Needed Extra Space
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println(" _   _      _        ______                ")
	fmt.Println("| | | |    | |       | ___ \\               ")
	fmt.Println("| |_| | ___| |_ __   | |_/ /_ _  __ _  ___ ")
	fmt.Println("|  _  |/ _ \\ | '_ \\  | __/  _  |/  _ |/  _ \\")
	fmt.Println("| | | |  __/ | |_) | | | | (_| |  (_||   __/")
	fmt.Println("\\_| |_/\\___|_| .__/  \\_|  \\__,_|\\__, |\\___|")
	fmt.Println("	     | |                 __/ |     ")
	fmt.Println("	     |_|                |___/      ")
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Flags:")
	fmt.Println("\t--warn (This flag is set to false by default. If the user wants to just warn and not fail the check set the warn flag to true.)")
	fmt.Println("\t--uBash (This flag is set to true by default. If the user wants to disable the bash analyzer they set the uBash flag to false.)")
	fmt.Println("\t--uMakefile (This flag is set to true by default. If the user wants to disable the makefile analyzer they set the uMakefile flag to false.)")
	fmt.Println("\t--uDockerfile (This flag is set to true by default. If the user wants to disable the dockerfile analyzer they set the uDockerfile flag to false.)")
	fmt.Println("\t--version (This flag can be used to check the current version of the project. Info is from the git tag.)")
	fmt.Println("\t--test (This flag is used to test the tool. Everytime code changes are made the test should pass.)")
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Examples:")
	fmt.Println("\t--warn: ./Soteria --warn=true ../ProjectName")
	fmt.Println("\t--uBash: ./Soteria --uBash=false ../ProjectName")
	fmt.Println("\t--uMakefile: ./Soteria --uMakefile=false ../ProjectName")
	fmt.Println("\t--uDockerfile: ./Soteria --uDockerfile=false ../ProjectName")
	fmt.Println("\t--version: ./Soteria --version")
	fmt.Println("\t--test: ./Soteria --test")
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("\n") // Needed Extra Space
}
