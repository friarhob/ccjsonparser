package main

import (
	"fmt"
	"os"

	"github.com/friarhob/ccjsonparser/internal/exitcodes"
	"github.com/friarhob/ccjsonparser/internal/parser"
)

func printHelpMessage() {
	fmt.Fprintln(os.Stderr, "Usage: ccjsonhelper [filepath]")
}

func execute() {
	if len(os.Args) > 2 {
		printHelpMessage()
		os.Exit(int(exitcodes.UsageError))
	}

	if len(os.Args) == 2 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		printHelpMessage()
		os.Exit(int(exitcodes.ValidJSON))
	}

	var file *os.File

	if len(os.Args) == 2 {
		var err error
		file, err = os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading file "+os.Args[1])
			os.Exit(int(exitcodes.ErrorReadingFile))
		}
		defer file.Close()
	} else {
		file = os.Stdin
	}

	if !parser.Validate(file) {
		fmt.Println("Invalid JSON")
		os.Exit(int(exitcodes.InvalidJSON))
	}

	fmt.Println("Valid JSON")
	os.Exit(int(exitcodes.ValidJSON))
}

func main() {
	execute()
}
