package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var moves = map[Vec]rune{{-1, 0}: '^', {0, 1}: '>', {1, 0}: 'v', {0, -1}: '<'}

type state struct {
	pos   Vec
	dir   Vec
	chLen int
}

type city struct {
	maze       [][]int
	chainRange []int
}

func (c *city) neighbours(cur state) []state {
	res := []state{}

	for m := range moves {
		if cur.dir.x*m.x < 0 || cur.dir.y*m.y < 0 {
			continue
		}

		if cur.dir == m && cur.chLen == c.chainRange[len(c.chainRange)-1] {
			continue
		}

		if cur.dir != m && cur.chLen < c.chainRange[0] && (cur.dir.x != 0 || cur.dir.y != 0) {
			continue
		}

		nPos := Vec{cur.pos.y + m.y, cur.pos.x + m.x}

		if !c.isInside(nPos) {
			continue
		}

		nChLen := cur.chLen + 1
		if cur.dir != m {
			nChLen = 1
		}
		res = append(res, state{pos: nPos, dir: m, chLen: nChLen})
	}

	return res
}

func (c *city) goalReached(cur state, goal state) bool {
	return cur.pos == goal.pos && cur.chLen >= c.chainRange[0]
}

func (c *city) costToMove(from state, to state) int {
	return c.maze[to.pos.y][to.pos.x]
}

func (c *city) distance(from state, to state) int {
	return 1
}

func (c *city) isInside(point Vec) bool {
	return 0 <= point.y && len(c.maze) > point.y && 0 <= point.x && len(c.maze[point.y]) > point.x
}

func main() {
	input := strings.Split(readFile("sample.txt"), "\n")
	maze := make([][]int, len(input)-1)
	for y, r := range input[:len(input)-1] {
		maze[y] = append(maze[y], strListToInt(r)...)
	}
	sY, sX := len(maze)-1, len(maze[len(maze)-1])-1

	// Part 1
	var c Dijkstras[state] = &city{maze: maze, chainRange: []int{0, 3}}
	// Part 2
	c = &city{maze: maze, chainRange: []int{4, 10}}

	path := shortestPath(c, state{pos: Vec{0, 0}}, state{pos: Vec{sY, sX}})
	path = append(path, state{pos: Vec{sY, sX}})

	printPath(maze, path)
	fmt.Fprintf(color.Output, "Total heatloss: %v\n", color.YellowString(strconv.Itoa(totalCost(maze, path))+"°C"))
}

func printPath(maze [][]int, path []state) {
	for y, r := range maze {
		for x, e := range r {
			isPath := false
			dir := Vec{}
			for _, p := range path {
				if p.pos.x == x && p.pos.y == y {
					isPath = true
					dir = p.dir
					break
				}
			}
			if y == len(maze)-1 && x == len(r)-1 {
				color.Set(color.FgYellow)
				fmt.Print("G")
			} else if y == 0 && x == 0 {
				color.Set(color.FgRed)
				fmt.Print("S")
			} else if isPath {
				color.Set(color.FgBlue)
				fmt.Print(string(moves[dir]))
			} else {
				fmt.Print(e)
			}
			color.Set(color.Reset)
		}
		fmt.Println()
	}
}

func totalCost(maze [][]int, path []state) int {
	total := 0
	for _, p := range path {
		total += maze[p.pos.y][p.pos.x]
	}
	return total
}

func strListToInt(list string) []int {
	c_list := []int{}
	for i := 0; i < len(list); i++ {
		num, err := strconv.ParseInt(string(list[i]), 10, 0)
		if err != nil {
			panic("Couldn't parse int")
		}
		c_list = append(c_list, int(num))
	}
	return c_list
}

func readFile(fn string) string {
	data, err := os.ReadFile(fn)

	if err != nil {
		panic("Couldn't read input file")
	}

	return string(data)
}
