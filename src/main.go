package main

import (
	"fmt"
	"os"

	"github.com/DomioKing653/FlowScript/src/lexer"
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

	for _, token := range tokens {
		token.Debug()
	}
}
