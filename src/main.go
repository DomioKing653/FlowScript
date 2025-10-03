package main

import (
	"fmt"
	"os"

	"github.com/DomioKing653/FlowScript/src/lexer"
	"github.com/DomioKing653/FlowScript/src/parser"
	"github.com/sanity-io/litter"
)

func main() {
	// recover from panics and print a red error with stack trace
	defer func() {
		if r := recover(); r != nil {
			printRed(fmt.Sprintf("%v", r))
			os.Exit(2)
		}
	}()

	file, err := os.ReadFile("./examples/prototype.flw")
	if err != nil {
		printRed(fmt.Sprintf("Error::IO->failed to read file: %v", err))
		os.Exit(-1)
	}
	tokens, err := lexer.Tokenize(string(file))
	if err != nil {
		printRed(err.Error())
		os.Exit(-1)
	}
	ast := parser.Parse(tokens)
	litter.Dump(ast)
}

// printRed writes the message to stderr in red (ANSI) and resets color.
func printRed(msg string) {
	red := "\x1b[31m"
	reset := "\x1b[0m"
	fmt.Fprintln(os.Stderr, red+msg+reset)
}
