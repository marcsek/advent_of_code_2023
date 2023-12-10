package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var pipes = map[string]string{
	"|": "NS",
	"-": "EW",
	"L": "NE",
	"J": "NW",
	"7": "SW",
	"F": "SE",
	"S": "NWSE",
}

var mvs = map[string][]int{
	"N": []int{-1, 0},
	"S": []int{1, 0},
	"E": []int{0, 1},
	"W": []int{0, -1},
}

func main() {
	input := strings.Split(readFile("test.txt"), "\n")
	maze, s_y, s_x := genMaze(input)
	traverse(s_y, s_x, 0, maze)

	y, x := findEmpty(maze)
	potCorrect := [][]int{}
	for y != -1 && x != -1 {
		found := [][]int{}
		if findAreas(y, x, maze, &found) {
			potCorrect = append(potCorrect, found...)
		}
		y, x = findEmpty(maze)
	}

	for _, pc := range potCorrect {
		y, x := pc[0], pc[1]
		maze[y][x] = "I"
	}

	total := 0
	for _, pc := range potCorrect {
		y, x := pc[0], pc[1]
		if y%2 == 0 && x%2 == 0 {
			total += 1
		}
	}
	printActual(maze)
	fmt.Println(total)
}

func findEmpty(maze [][]string) (int, int) {
	for y, r := range maze {
		for x, c := range r {
			if c != "O" && c != "!" && c != "S" {
				return y, x
			}
		}
	}
	return -1, -1
}

func findAreas(y, x int, maze [][]string, found *[][]int) bool {
	isInside := true
	if y < 0 || y >= len(maze) || x < 0 || x >= len(maze[0]) {
		return false
	} else if maze[y][x] == "!" || maze[y][x] == "O" || maze[y][x] == "S" {
		return true
	}

	maze[y][x] = "O"
	*found = append(*found, []int{y, x})
	for _, r := range "NSEW" {
		n_y, n_x := mvs[string(r)][0], mvs[string(r)][1]
		if !findAreas(y+n_y, x+n_x, maze, found) {
			isInside = false
		}
	}
	return isInside
}

func traverse(y, x, m int, maze [][]string) int {
	cur := maze[y][x]
	dirs := pipes[cur]
	moves := [][]int{}
	all_len := []int{}

	for _, d := range dirs {
		moves = append(moves, mvs[string(d)])
	}

	for i, move := range moves {
		n_y, n_x := move[0], move[1]
		if y+n_y < 0 || y+n_y >= len(maze) || x+n_x < 0 || x+n_x >= len(maze[0]) {
			continue
		}
		if next := maze[y+n_y][x+n_x]; next != "." && next[0] != '!' {
			if next == "S" && m > 3 {
				return m
			}
			nextDir := pipes[next]
			dir := pipes[cur][i]

			if checkIfValid(string(dir), nextDir) {
				if cur != "S" {
					maze[y][x] = "!"
				}
				all_len = append(all_len, traverse(n_y+y, n_x+x, m+1, maze))
			}
		} else if next := maze[y+n_y][x+n_x]; next != "." {
			if cur != "S" {
				maze[y][x] = "!"
			}
		}
	}
	if len(all_len) != 0 {
		return slices.Max(all_len)
	}
	return m
}

func checkIfValid(md, ndir string) bool {
	m := mvs[md]
	for _, c := range ndir {
		t := mvs[string(c)]
		if y, x := m[0]+t[0], m[1]+t[1]; y == 0 && x == 0 {
			return true
		}
	}

	return false
}

func genMaze(input []string) ([][]string, int, int) {
	maze := [][]string{}
	y, x := 0, 0
	for i_y, line := range input[:len(input)-1] {
		row := []string{}
		for i_x, c := range line {
			if c == 'S' {
				y, x = i_y*2, i_x*2
			}
			toAdd := "."
			if c == 'F' || c == 'L' || c == '-' || c == 'S' {
				toAdd = "-"
			}
			row = append(row, string(c), toAdd)
		}
		e_row := []string{}
		for i := 0; i < len(row); i++ {
			toAdd := "."
			if row[i] == "|" || row[i] == "S" || row[i] == "7" || row[i] == "F" {
				toAdd = "|"
			}
			e_row = append(e_row, toAdd)
		}
		maze = append(maze, row, e_row)
	}
	return maze, y, x
}

func printMaze(maze [][]string) {
	for _, r := range maze {
		for _, c := range r {
			fmt.Print(c)
		}
		fmt.Println("")
	}
}
func printActual(maze [][]string) {
	for y := 0; y < len(maze); y += 2 {
		for x := 0; x < len(maze[y]); x += 2 {
			fmt.Print(maze[y][x])
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
