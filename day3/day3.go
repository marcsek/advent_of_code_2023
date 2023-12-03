package main

import (
	"Strings"
	"fmt"
	"slices"
	"strconv"
	"unicode"
)

var not_symbols = []string{
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	".",
}

func main() {
	dirs := [8][2]int{
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
	}

	rows := [][]string{}
	for _, r := range strings.Split(test, "\n") {
		rows = append(rows, strings.Split(r, ""))
	}

	s := 0
	for y, r := range rows {
		for x, char := range rows[y] {
			if !slices.Contains(not_symbols, char) {
				for _, dir := range dirs {
					n_y, n_x := y+dir[0], x+dir[1]
					if 0 <= n_y && n_y < len(rows) && 0 <= n_x && n_x < len(r) {
						if unicode.IsDigit(rune(rows[n_y][n_x][0])) {
							num := makeNumber(&rows[n_y], n_x)
							s += num
							fmt.Println(num)
						}
					}
				}
			}
		}
	}

	fmt.Println(s)
}

func makeNumber(r *[]string, x int) int {
	num := (*r)[x]

	p_l := x - 1
	p_r := x + 1

	for p_l >= 0 && p_r < len(*r) {
		d1, d2 := false, false
		if unicode.IsDigit(rune((*r)[p_l][0])) {
			num = (*r)[p_l] + num
			(*r)[p_l] = "."
			p_l--
			d1 = true
		}
		if unicode.IsDigit(rune((*r)[p_r][0])) {
			num += (*r)[p_r]
			(*r)[p_r] = "."
			p_r++
			d2 = true
		}
		if !d1 && !d2 {
			break
		}
	}

	n, err := strconv.ParseInt(num, 10, 0)

	if err != nil {
		panic("Zle je")
	}

	return int(n)
}

var test = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
