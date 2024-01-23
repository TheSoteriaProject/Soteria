package main;

import (
	"fmt"
	"os"
)

func help() {
	fmt.Println("Soteria is a linter that is used for...");
	fmt.Println("\nFlags:");
	fmt.Println("\t --Warn");
}

func version() {
	fmt.Println("Version 0.0.1");
}

func run(flag string) {
	fmt.Println("Running...");
	
	if flag == "--warn" {
		fmt.Println("Warning Mode: On");
	} else {
		fmt.Println("Warning Mode: Off");
	}
}

func main() {
	if len(os.Args) > 1{
		arg := os.Args[1];
		switch arg {
			case "--help", "help":
				help();
			case "--version", "version":
				version();
			case "--warn":
				run(arg);
			default:
				fmt.Println("Invalid CLI argument.\nTry --help");
		}
	} else {
		run("NULL");
	}
}
