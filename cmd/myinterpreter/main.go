package main

import (
	"bytes"
	"fmt"
	"os"

	loxinterpreter "github.com/codecrafters-io/interpreter-starter-go/internal/lox-interpreter"
	"github.com/codecrafters-io/interpreter-starter-go/internal/scanner"
)

func run(content []byte) {
	if len(content) > 0 {
		s := scanner.NewScanner(bytes.Runes(content))
		tokens := s.ScanTokens()
		for _, t := range tokens {
			fmt.Fprintln(os.Stdout, t)
		}
	} else {
		fmt.Println("EOF  null")
	}
}

func runFile() {
	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	run(fileContents)
	if loxinterpreter.HadError {
		os.Exit(65)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	runFile()
}
