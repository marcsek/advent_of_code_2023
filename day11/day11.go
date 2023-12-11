package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const multBy = 1_000_000

func main() {
	input := strings.Split(readFile("test.txt"), "\n")
	expans := parseInput(input[:len(input)-1])
	edges := findEdges(input)

	fmt.Println(findLength(edges, expans))
}

func findLength(edges [][]int, expans [][]int) int {
	total := 0
	for a1 := 0; a1 < len(edges); a1++ {
		for a2 := a1 + 1; a2 < len(edges); a2++ {
			c1, c2 := edges[a1], edges[a2]
			y1, x1, y2, x2 := c1[0], c1[1], c2[0], c2[1]

			mult := 0
			for _, expan := range expans {
				ey, ex := expan[0], expan[1]
				y1, y2 := int(math.Min(float64(y1), float64(y2))), int(math.Max(float64(y1), float64(y2)))
				x1, x2 := int(math.Min(float64(x1), float64(x2))), int(math.Max(float64(x1), float64(x2)))

				if y1 < ey && y2 > ey {
					mult += multBy - 1
				}
				if x1 < ex && x2 > ex {
					mult += multBy - 1
				}
			}
			total += int(math.Abs(float64((x1 - x2)))) + int(math.Abs(float64((y1 - y2)))) + mult
		}
	}
	return total
}

func findEdges(maze []string) [][]int {
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

func parseInput(input []string) [][]int {
	expan := [][]int{}
	for y, r := range input {
		hasGalaxy := false
		for _, s := range r {
			if s == '#' {
				hasGalaxy = true
			}
		}
		if !hasGalaxy {
			expan = append(expan, []int{y, 0})
		}
	}
	for x := range input[0] {
		hasGalaxy := false
		for y := range input {
			if input[y][x] == '#' {
				hasGalaxy = true
			}
		}
		if !hasGalaxy {
			expan = append(expan, []int{0, x})
		}
	}
	return expan
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
