package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := strings.Split(readFile("sample.txt"), "\n")
	maze := parseInput(input[:len(input)-1])

	printMaze(maze)
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
	skip := false
	for x := range maze[0] {
		hasGalaxy := false
		for y := range maze {
			if maze[y][x] == '#' {
				hasGalaxy = true
			}
		}
		if !hasGalaxy && !skip {
			for y_n := range maze {
				maze[y_n] = append(maze[y_n][:x+1], maze[y_n][x:]...)
				maze[y_n][x] = '.'
			}
			skip = true
		} else {
			skip = false
		}
	}

	return maze
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
