package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/internal/scanner"
)

func runFile() {

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

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(fileContents) > 0 {
		s := scanner.NewScanner(string(fileContents))
		tokens,err := s.ScanTokens()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, t := range(tokens) {
			fmt.Fprintln(os.Stdout, t)
		}
	} else {
		fmt.Println("EOF  null")
	}
}
