package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type dig struct {
	dir string
	num int
}

type vec struct {
	y int
	x int
}

func main() {
	input := strings.Split(readFile("test.txt"), "\n")
	digs := parseInput(input[:len(input)-1])
	coords, border := generateCoords(digs)

	area := getAreaWithCoords(coords)
	area = (int(math.Abs(float64(area))) + border + 1) / 2
	fmt.Printf("result: %v\n", area)
}

func getAreaWithCoords(coords []vec) int {
	area := 0
	for i := 1; i < len(coords); i++ {
		v1 := coords[i-1]
		v2 := coords[i]

		area += (v1.y + v2.y) * (v1.x - v2.x)
	}
	return area
}

func generateCoords(digs []dig) ([]vec, int) {
	py, px := 0, 0
	coords := []vec{}
	border := 1

	for _, d := range digs {
		if d.dir == "R" {
			c2 := vec{py, px + d.num}
			coords = append(coords, c2)
			py, px = c2.y, c2.x
		} else if d.dir == "L" {
			c1 := vec{py, px - d.num}
			coords = append(coords, c1)
			py, px = c1.y, c1.x
		} else if d.dir == "D" {
			c2 := vec{py + d.num, px}
			coords = append(coords, c2)
			py, px = c2.y, c2.x
		} else {
			c1 := vec{py - d.num, px}
			coords = append(coords, c1)
			py, px = c1.y, c1.x
		}
		border += d.num
	}
	return coords, border
}

func parseInput(input []string) []dig {
	digs := make([]dig, len(input))
	dirs := []rune{'R', 'D', 'L', 'U'}
	for i, line := range input {
		sp := strings.Split(line, " ")
		hex := sp[2][2 : len(sp[2])-1]
		ld := rune(hex[5])
		dir := dirs[decToInt(string(ld))]
		num := hexToInt(hex[:5])

		digs[i] = dig{string(dir), num}
	}
	return digs
}

func hexToInt(i string) int {
	in, err := strconv.ParseInt(i, 16, 0)

	if err != nil {
		panic("aja")
	}
	return int(in)
}

func decToInt(i string) int {
	in, err := strconv.ParseInt(i, 10, 0)

	if err != nil {
		panic("aja")
	}
	return int(in)
}

func readFile(fn string) string {
	data, err := os.ReadFile(fn)

	if err != nil {
		panic("Couldn't read input file")
	}

	return string(data)
}
