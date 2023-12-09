package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	input := strings.Split(readFile("test.txt"), "\n")
	dirs := ""
	sepIdx := 0

	for i, line := range input {
		if line == "" {
			sepIdx = i
			break
		}
		dirs += line
	}

	inst := map[string][2]string{}
	startNodes := []string{}
	for _, line := range input[sepIdx+1 : len(input)-1] {
		spl := strings.Split(line, " = ")
		p := strings.Split(spl[1], ", ")
		id := spl[0]
		paths := [2]string{p[0][1:], p[1][:3]}

		if id[2] == 'A' {
			startNodes = append(startNodes, id)
		}
		inst[id] = paths
	}

	idx := 0
	firstZs := []int{}
	for len(firstZs) != len(startNodes) {
		for i := range startNodes {
			if dirs[idx%len(dirs)] == 'R' {
				startNodes[i] = inst[startNodes[i]][1]
			} else {
				startNodes[i] = inst[startNodes[i]][0]
			}
			if startNodes[i][2] == 'Z' {
				firstZs = append(firstZs, idx+1)
			}
		}
		idx++
	}

	fmt.Printf("result: %v\n", LCM(firstZs[0], firstZs[1], firstZs...))
	fmt.Printf("total time: %v", time.Since(start))
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}

func readFile(fn string) string {
	data, err := os.ReadFile(fn)

	if err != nil {
		panic("Couldn't read input file")
	}

	return string(data)
}
