package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := strings.Split(readFile("sample.txt"), "\n")
	maze := parseInput(input[:len(input)-1])

	fmt.Println(len(findEdges(maze)))
}

func findLength(maze [][]rune) int {
	for a1 := 0; a1 < len(maze); a1++ {
		for a2 := a1 + 1; a2 < len(maze); a2++ {

		}
	}
}

func findEdges(maze [][]rune) [][]int {
	edges := [][]int{}

	for y, r := range maze {
		for x, c := range r {
			if c == '#' {
				edges = append(edges, []int{y, x})
			}
		}
	}
	return edges
}

func parseInput(input []string) [][]rune {
	maze := [][]rune{}
	for _, r := range input {
		row := []rune{}
		hasGalaxy := false
		for _, s := range r {
			if s == '#' {
				hasGalaxy = true
			}
			row = append(row, s)
		}
		maze = append(maze, row)
		if !hasGalaxy {
			empty := []rune{}
			for i := 0; i < len(row); i++ {
				empty = append(empty, '.')
			}
			maze = append(maze, empty)
		}
	}
	correct := make([][]rune, len(maze))
	for x := range maze[0] {
		hasGalaxy := false
		for i := 0; i < len(maze); i++ {
			if maze[i][x] == '#' {
				hasGalaxy = true
			}
		}
		for y := range maze {
			correct[y] = append(correct[y], maze[y][x])
			if !hasGalaxy {
				correct[y] = append(correct[y], '.')
			}
		}
	}
	return correct
}

func printMaze(maze [][]rune) {
	for _, r := range maze {
		for _, c := range r {
			fmt.Print(string(c))
		}
		fmt.Println("")
	}
}

func readFile(fn string) string {
	data, err := os.ReadFile(fn)

	if err != nil {
		panic("Couldn't read input file")
	}

	return string(data)
}
