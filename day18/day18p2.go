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

// 2685885422866
//  952408144115

func main() {
	input := strings.Split(readFile("sample.txt"), "\n")
	//digs := parseInput(input[:len(input)-1])
	digs := parseInputOld(input[:len(input)-1])
	y1, x1, y2, x2 := getHoleSize(digs)
	//sy, sx := int(math.Abs(float64(y1)))+y2, int(math.Abs(float64(x1)))+x2
	py, px := int(math.Abs(float64(y1))), int(math.Abs(float64(x1)))
	walls := generateWalls(py, px, digs)
	fmt.Println(walls)

	area := getAreaWithWalls(walls)

	fmt.Println(area)
	fmt.Println(y1, x1, y2, x2)
}

func getAreaWithWalls(walls [][]vec) int {
	area := 0
	cmprd := [][]int{}
	overlaps := [][]int{}
	for i := 0; i < len(walls)-1; i++ {
		for j := i + 1; j < len(walls); j++ {
			alreadyDid := false
			for _, c := range cmprd {
				pi, pj := c[1], c[0]
				if pi == i && pj == j {
					alreadyDid = true
				}
			}
			if alreadyDid {
				continue
			}
			w1, w2 := walls[i], walls[j]
			x1, x2 := w1[0].x, w2[0].x
			w1y1, w1y2, w2y1, w2y2 := w1[0].y, w1[1].y, w2[0].y, w2[1].y
			rs := int(math.Max(float64(w1y1), float64(w2y1)))
			re := int(math.Min(float64(w1y2), float64(w2y2)))
			cmprd = append(cmprd, []int{i, j})
			if w1y2 > w2y1 && w1y1 < w2y2 {
				rng := int(math.Abs(float64(re-rs))) + 1
				w := int(math.Abs(float64(x2-x1))) + 1
				for _, ov := range overlaps {
					l := ov[0]
					if re == l {
						fmt.Println(re, w)
						area -= w
						break
					}
				}
				fmt.Println(rs, re, w1, w2, i, j, w*rng)
				area += w * rng
				overlaps = append(overlaps, []int{rs, x1, x2}, []int{re, x1, x2})
			}
		}
	}
	return area
}

func generateWalls(sy, sx int, digs []dig) [][]vec {
	py, px := sy, sx
	walls := [][]vec{}

	for _, d := range digs {
		if d.dir == "R" {
			//c1 := vec{py, px}
			c2 := vec{py, px + d.num}
			py, px = c2.y, c2.x
		} else if d.dir == "L" {
			//c2 := vec{py, px}
			c1 := vec{py, px - d.num}
			py, px = c1.y, c1.x
		} else if d.dir == "D" {
			c1 := vec{py, px}
			c2 := vec{py + d.num, px}
			walls = append(walls, []vec{c1, c2})
			py, px = c2.y, c2.x
		} else {
			c2 := vec{py, px}
			c1 := vec{py - d.num, px}
			walls = append(walls, []vec{c1, c2})
			py, px = c1.y, c1.x
		}
	}
	return walls
}

func parseInputOld(input []string) []dig {
	digs := make([]dig, len(input))
	for i, line := range input {
		sp := strings.Split(line, " ")
		dir, num := sp[0], decToInt(sp[1])
		digs[i] = dig{dir, num}
	}
	return digs
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

func getHoleSize(digs []dig) (int, int, int, int) {
	y, x := 0, 0
	y1, x1, y2, x2 := 0, 0, 0, 0

	for _, d := range digs {
		if d.dir == "R" {
			x += d.num
		} else if d.dir == "L" {
			x -= d.num
		} else if d.dir == "D" {
			y += d.num
		} else {
			y -= d.num
		}
		if y1 > y {
			y1 = y
		}
		if y2 < y {
			y2 = y
		}
		if x1 > x {
			x1 = x
		}
		if x2 < x {
			x2 = x
		}
	}
	return y1, x1, y2, x2
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
