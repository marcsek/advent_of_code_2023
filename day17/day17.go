package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type vec struct {
	y int
	x int
}

var moves = []vec{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func main() {
	input := strings.Split(readFile("test.txt"), "\n")
	maze := make([][]int, len(input)-1)
	for y, r := range input[:len(input)-1] {
		maze[y] = append(maze[y], strListToInt(r)...)
	}
	total := 0
	min := new(int)
	*min = int(math.Inf(1))
	lastPos := vec{0, 0}
	i := 0
	for true {
		winPos := new(vec)
		walkMaze(maze, min, 0, lastPos, vec{0, -1}, 0, []vec{}, 15, winPos)
		total += *min
		lastPos = *winPos
		if lastPos.x == len(maze[0])-1 && lastPos.y == len(maze)-1 {
			break
		}
		*min = int(math.Inf(1))
		if i > 100 {
			break
		}
		i++
	}
}

func walkMaze(
	maze [][]int,
	min *int,
	stSteps int,
	pos vec,
	pDir vec,
	total int,
	pp []vec,
	buff int,
	winPos *vec,
) {
	if total >= *min || pos.y < 0 || pos.y >= len(maze) || pos.x < 0 || pos.x >= len(maze[0]) {
		return
	}

	if buff <= 0 || (pos.y == len(maze)-1 && pos.x == len(maze[0])-1) {
		if *min > total {
			*winPos = pos
			*min = total
			//showPath(pp, maze)
		}
		return
	}

	for _, p := range pp {
		if p.y == pos.y && p.x == pos.x {
			return
		}
	}

	pp = append(pp, pos)
	cNum := maze[pos.y][pos.x]
	for _, move := range moves {
		if move.x != -pDir.x || move.y != -pDir.y {
			newPos := vec{pos.y + move.y, pos.x + move.x}
			if move.x == pDir.x && move.y == pDir.y {
				if stSteps < 2 {
					walkMaze(maze, min, stSteps+1, newPos, move, total+cNum, pp, buff-1, winPos)
				}
			} else {
				walkMaze(maze, min, 0, newPos, move, total+cNum, pp, buff-1, winPos)
			}
		}
	}
}

func showPath(path []vec, maze [][]int) {
	sum := 0
	for y, r := range maze {
		for x, e := range r {
			find := false
			for _, p := range path {
				if p.y == y && p.x == x {
					find = true
					break
				}
			}
			if find {
				fmt.Print("#")
				sum += e
			} else {
				fmt.Print(e)
			}
		}
		fmt.Println("")
	}
	fmt.Println(sum)
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
