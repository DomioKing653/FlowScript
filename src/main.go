package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, err := os.ReadFile("./examples/math.flw")
	if err != nil {
		panic(err)
	}
	source := string(bytes)

	fmt.Println(source)
}
