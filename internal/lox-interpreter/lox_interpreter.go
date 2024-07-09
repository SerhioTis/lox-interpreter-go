package loxinterpreter

import (
	"fmt"
	"os"
)

var HadError = false

func report(line int, where string, message string) {
	fmt.Fprintf(os.Stderr, "[line %v] Error%s: %s\n", line, where, message)
	HadError = true
}

func Error(line int, message string) {
	report(line, "", message)
}
