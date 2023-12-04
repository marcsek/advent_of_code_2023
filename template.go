package main

import (
	"fmt"
	"os"
)

func main() {
	input := readFile("daco.txt")

	fmt.Println(input)
}

func readFile(fn string) string {
	data, err := os.ReadFile(fn)

	if err != nil {
		panic("Couldn't read input file")
	}

	return string(data)
}
