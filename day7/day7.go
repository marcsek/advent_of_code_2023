package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

var strengths = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

func main() {
	input := readFile("sample.txt")
	lines := strings.Split(input, "\n")
	cards := [][]string{}
	values := [7][][]string{}
	sum := 0

	for _, line := range lines[:len(lines)-1] {
		cards = append(cards, strings.Split(line, " "))
	}

	for _, card := range cards {
		mp := make(map[string]int)
		for _, char := range card[0] {
			_, ok := mp[string(char)]
			if !ok {
				mp[string(char)] = 1
			} else {
				mp[string(char)] += 1
			}
		}
		mx := slices.Max(maps.Values(mp))
		if len(mp) == 1 {
			values[0] = append(values[0], card)
		} else if len(mp) == 2 {
			if mx == 4 {
				values[1] = append(values[1], card)
			} else {
				values[2] = append(values[2], card)
			}
		} else if len(mp) == 3 {
			mx := slices.Max(maps.Values(mp))
			if mx == 3 {
				values[3] = append(values[3], card)
			} else {
				values[4] = append(values[4], card)
			}
		} else if len(mp) == 4 {
			values[5] = append(values[5], card)
		} else {
			values[6] = append(values[6], card)
		}
	}
	r := 1
	slices.Reverse(values[:])
	for _, value := range values {
		if len(value) > 1 {
			fmt.Println(compareCards(value))
			r += 1
		} else if len(value) == 1 {
			v := toInt(value[0][1])
			sum += v * r
			r += 1
		} else {
			r += 1
		}
	}
	fmt.Println(sum)
}

func compareCards(cards [][]string) [][]string {
	bv := [][]int{}
	vys := [][]string{}
	for i, pair := range cards {
		card := pair[0]
		s := 0
		for _, char := range card {
			s += strengths[string(char)]
		}
		bv = append(bv, []int{s, i})
	}
	sort.SliceStable(bv, func(i, j int) bool {
		return bv[i][0] < bv[j][0]
	})
	for _, b := range bv {
		vys = append(vys, cards[b[1]])
	}
	return vys
}

func toInt(i string) int {
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
