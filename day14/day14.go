package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var cycles = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func main() {
	input := strings.Split(readFile("test.txt"), "\n")
	maze := generateMaze(input[:len(input)-1])

	pv := [][][]rune{}
	cs := []int{}
	for i := 0; i < 200; i++ {
		found := false
		t := makeCycle(maze)
		cs = append(cs, t)
		for pi, p := range pv {
			allEq := true
			for y := range p {
				if !slices.Equal(p[y], maze[y]) {
					allEq = false
				}
			}
			if allEq {
				found = true
				fmt.Println(cs[pi : len(cs)-1][(1000000000-pi-1)%(len(cs)-pi-1)])
				break
			}
		}
		if found {
			break
		}
		pv = append(pv, cloneTwoDim(maze))
	}
}

func makeCycle(maze [][]rune) int {
	tTotal := 0
	for _, cycle := range cycles {
		total := 0
		cY, cX := cycle[0], cycle[1]
		if cY != 0 {
			for y := 0; y < len(maze); y++ {
				for x := 0; x < len(maze[y]); x++ {
					aY := y
					if cY == 1 {
						aY = len(maze) - y - 1
					}
					if maze[aY][x] == 'O' {
						nY, nX := makeFall(aY, x, cycle, maze)
						maze[aY][x] = '.'
						maze[nY][nX] = 'O'
					}
				}
			}
		} else {
			for x := 0; x < len(maze[0]); x++ {
				for y := 0; y < len(maze); y++ {
					aX := x
					if cX == 1 {
						aX = len(maze[0]) - x - 1
					}
					if maze[y][aX] == 'O' {
						nY, nX := makeFall(y, aX, cycle, maze)
						maze[y][aX] = '.'
						maze[nY][nX] = 'O'
						total += len(maze) - nY
					}
				}
			}
		}
		tTotal = total
	}
	return tTotal
}

func makeFall(y, x int, dir []int, maze [][]rune) (int, int) {
	pY, pX := y, x
	nY, nX := y+dir[0], x+dir[1]

	for true {
		bnds := nY >= 0 && nY < len(maze) && nX >= 0 && nX < len(maze[0])
		if !bnds {
			return pY, pX
		}
		corCh := maze[nY][nX] != '#' && maze[nY][nX] != 'O'
		if !corCh {
			return pY, pX
		}
		pY, pX = nY, nX
		nY, nX = nY+dir[0], nX+dir[1]
	}
	return nY, nX
}

func cloneTwoDim(i [][]rune) [][]rune {
	n := [][]rune{}

	for _, r := range i {
		n = append(n, slices.Clone(r))
	}
	return n
}

func generateMaze(input []string) [][]rune {
	maze := make([][]rune, 0, len(input))

	for _, row := range input {
		r := make([]rune, 0, len(row))
		for _, ch := range row {
			r = append(r, ch)
		}
		maze = append(maze, r)
	}
	return maze
}

func printMaze(maze [][]rune) {
	for _, row := range maze {
		fmt.Println(string(row))
	}
}

func readFile(fn string) string {
	data, err := os.ReadFile(fn)

	if err != nil {
		panic("Couldn't read input file")
	}

	return string(data)
}
