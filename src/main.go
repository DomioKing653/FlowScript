package main

import (
	"os"

	"github.com/DomioKing653/FlowScript/src/lexer"
)

func main() {
	file, err := os.ReadFile("./examples/math.flw")
	if err != nil {
		panic(err)
	}
	tokens := lexer.Tokenize(string(file))

	for _, token := range tokens {
		token.Debug()
	}
}
