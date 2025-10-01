package main

import (
	"fmt"
	"os"

	"github.com/DomioKing653/FlowScript/src/lexer"
	"github.com/DomioKing653/FlowScript/src/parser"
)

func main() {
	file, err := os.ReadFile("./examples/main.flw")
	if err != nil {
		panic(err)
	}
	tokens, err := lexer.Tokenize(string(file))
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	ast := parser.Parse(tokens)
	for _, token := range tokens {
		token.Debug()
	}
	fmt.Println(ast)
}
