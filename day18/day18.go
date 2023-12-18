//package main
//
//import (
//	"fmt"
//	"math"
//	"os"
//	"strconv"
//	"strings"
//)
//
//type dig struct {
//	dir   string
//	num   int
//	color string
//}
//
//type vec struct {
//	y int
//	x int
//}
//
//func main() {
//	input := strings.Split(readFile("sample.txt"), "\n")
//	maze := [][]rune{}
//	digs := parseInput(input[:len(input)-1])
//	y1, x1, y2, x2 := getHoleSize(digs)
//	sy, sx := int(math.Abs(float64(y1-1)))+y2+1, int(math.Abs(float64(x1-1)))+x2+1
//	py, px := int(math.Abs(float64(y1-1))), int(math.Abs(float64(x1-1)))
//	coords := generateCoords(py, px, digs)
//
//	for y := 0; y <= sy; y++ {
//		r := []rune{}
//		for x := 0; x <= sx; x++ {
//			didFind := false
//			for _, c := range coords {
//				if y == c.y && x == c.x {
//					didFind = true
//				}
//			}
//			if didFind {
//				r = append(r, '#')
//			} else {
//				r = append(r, '.')
//			}
//		}
//		maze = append(maze, r)
//	}
//
//	total := floodFill(0, 0, 0, maze)
//	//printMaze(maze)
//	//fmt.Println(y1, x1, y2, x2, sy, sx)
//	fmt.Printf("result: %v\n", ((sy-1)*(sx-1))-(total-(2*sy)-(2*sx)))
//}
//
//func floodFill(y, x, total int, maze [][]rune) int {
//	sy, sx := len(maze), len(maze[0])
//	if 0 > y || y >= sy || 0 > x || x >= sx || maze[y][x] == '#' || maze[y][x] == '!' {
//		return total
//	}
//	maze[y][x] = '!'
//
//	total += floodFill(y+1, x, 0, maze)
//	total += floodFill(y-1, x, 0, maze)
//	total += floodFill(y, x+1, 0, maze)
//	total += floodFill(y, x-1, 0, maze)
//	return total + 1
//}
//
//func parseInput(input []string) []dig {
//	digs := make([]dig, len(input))
//	for i, line := range input {
//		sp := strings.Split(line, " ")
//		dir, num, color := sp[0], toInt(sp[1]), sp[2]
//		digs[i] = dig{dir, num, color}
//	}
//	return digs
//}
//
//func generateCoords(sy, sx int, digs []dig) []vec {
//	cy, cx := sy, sx
//	pos := []vec{}
//
//	for _, d := range digs {
//		if d.dir == "R" {
//			for i := cx; i < cx+d.num; i++ {
//				pos = append(pos, vec{cy, i})
//			}
//			cx += d.num
//		} else if d.dir == "L" {
//			for i := cx; i >= cx-d.num; i-- {
//				pos = append(pos, vec{cy, i})
//			}
//			cx -= d.num
//		} else if d.dir == "D" {
//			for i := cy; i < cy+d.num; i++ {
//				pos = append(pos, vec{i, cx})
//			}
//			cy += d.num
//		} else {
//			for i := cy; i >= cy-d.num; i-- {
//				pos = append(pos, vec{i, cx})
//			}
//			cy -= d.num
//		}
//	}
//	return pos
//}
//
//func getHoleSize(digs []dig) (int, int, int, int) {
//	y, x := 0, 0
//	y1, x1, y2, x2 := 0, 0, 0, 0
//	//pos := []vec{}
//
//	for _, d := range digs {
//		if d.dir == "R" {
//			//for i := x; i < x+d.num; i++ {
//			//	pos = append(pos, vec{y, i})
//			//}
//			x += d.num
//		} else if d.dir == "L" {
//			//for i := x; i >= x-d.num; i-- {
//			//	pos = append(pos, vec{y, i})
//			//}
//			x -= d.num
//		} else if d.dir == "D" {
//			//for i := y; i < y+d.num; i++ {
//			//	pos = append(pos, vec{i, x})
//			//}
//			y += d.num
//		} else {
//			//for i := y; i >= y-d.num; i-- {
//			//	pos = append(pos, vec{i, x})
//			//}
//			y -= d.num
//		}
//		if y1 > y {
//			y1 = y
//		}
//		if y2 < y {
//			y2 = y
//		}
//		if x1 > x {
//			x1 = x
//		}
//		if x2 < x {
//			x2 = x
//		}
//	}
//	return y1, x1, y2, x2
//}
//
//func printMaze(maze [][]rune) {
//	for _, r := range maze {
//		fmt.Println(string(r))
//	}
//}
//
//func toInt(i string) int {
//	in, err := strconv.ParseInt(i, 10, 0)
//
//	if err != nil {
//		panic("aja")
//	}
//	return int(in)
//}
//
//func readFile(fn string) string {
//	data, err := os.ReadFile(fn)
//
//	if err != nil {
//		panic("Couldn't read input file")
//	}
//
//	return string(data)
//}
