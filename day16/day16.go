package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type vec struct {
	y int
	x int
}

type Beam struct {
	pos vec
	dir vec
}

func main() {
	input := strings.Split(readFile("test.txt"), "\n")
	maze := [][]rune{}
	for _, r := range input[:len(input)-1] {
		maze = append(maze, []rune(r))
	}
	beams := []Beam{}
	max := 0
	for y := range maze {
		beams = append(beams, Beam{vec{y, 0}, vec{0, 1}})
		beams = append(beams, Beam{vec{y, len(maze[0]) - 1}, vec{0, -1}})
	}
	for x := range maze[0] {
		beams = append(beams, Beam{vec{0, x}, vec{1, 0}})
		beams = append(beams, Beam{vec{len(maze) - 1, x}, vec{-1, 0}})
	}

	for _, b := range beams {
		if cur := startBeam(b, maze); cur > max {
			max = cur
		}
	}
	fmt.Println(max)
}

func startBeam(beam Beam, maze [][]rune) int {
	beams := []Beam{beam}
	passMap := initPassMap(len(maze), len(maze[0]))

	i := 0
	for len(beams) > 0 {
		idx := i % len(beams)
		beam := &(beams[idx])
		isOut, newBeam := beam.makeMove(maze, &passMap)
		if isOut {
			beams = slices.Delete(beams, idx, idx+1)
		}
		if newBeam != nil {
			beams = append(beams, *newBeam)
		}
		i++
	}
	return countEnergized(passMap)
}

func (beam *Beam) isInCycle(passMap *[][][]vec) bool {
	y, x := beam.pos.y, beam.pos.x
	if 0 > y || y >= len(*passMap) || 0 > x || x >= len((*passMap)[0]) {
		return true
	}

	for _, pass := range (*passMap)[beam.pos.y][beam.pos.x] {
		if pass.y == beam.dir.y && pass.x == beam.dir.x {
			return true
		}
	}
	(*passMap)[y][x] = append((*passMap)[y][x], vec{beam.dir.y, beam.dir.x})
	return false
}

func (beam *Beam) makeMove(maze [][]rune, passMap *[][][]vec) (bool, *Beam) {
	y, x := beam.pos.y, beam.pos.x
	if 0 > y || y >= len(maze) || 0 > x || x >= len((maze)[0]) || beam.isInCycle(passMap) {
		return true, nil
	}

	tile := maze[y][x]
	var newBeam *Beam = nil

	if tile == '/' {
		if beam.dir.x != 0 {
			beam.dir.y = -beam.dir.x
			beam.dir.x = 0
		} else if beam.dir.y != 0 {
			beam.dir.x = -beam.dir.y
			beam.dir.y = 0
		}
	} else if tile == '\\' {
		if beam.dir.x != 0 {
			beam.dir.y = beam.dir.x
			beam.dir.x = 0
		} else if beam.dir.y != 0 {
			beam.dir.x = beam.dir.y
			beam.dir.y = 0
		}
	} else if tile == '|' {
		if beam.dir.x != 0 {
			beam.dir.x = 0
			beam.dir.y = -1
			newBeam = &Beam{vec{y, x}, vec{1, 0}}
		}
	} else if tile == '-' {
		if beam.dir.y != 0 {
			beam.dir.y = 0
			beam.dir.x = 1
			newBeam = &Beam{vec{y, x}, vec{0, -1}}
		}
	}
	beam.pos.y, beam.pos.x = y+beam.dir.y, x+beam.dir.x
	return false, newBeam
}

func initPassMap(sy, sx int) [][][]vec {
	passMap := make([][][]vec, sy)
	for y := 0; y < sy; y++ {
		for x := 0; x < sx; x++ {
			passMap[y] = make([][]vec, sx)
		}
	}
	return passMap
}

func printMaze(maze [][]rune) {
	for _, r := range maze {
		fmt.Println(string(r))
	}
}

func countEnergized(passMap [][][]vec) int {
	total := 0
	for _, r := range passMap {
		for _, t := range r {
			if len(t) != 0 {
				total++
			}
		}
	}
	return total
}

func cloneTwoDim(i [][]rune) [][]rune {
	n := [][]rune{}

	for _, r := range i {
		n = append(n, slices.Clone(r))
	}
	return n
}

func readFile(fn string) string {
	data, err := os.ReadFile(fn)

	if err != nil {
		panic("Couldn't read input file")
	}

	return string(data)
}
