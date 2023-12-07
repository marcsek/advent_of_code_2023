package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

var strengths = map[string]int{
	"J": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
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
		c, hasJ := mp["J"]
		plusTier := 0
		if hasJ {
			plusTier = 1
		}
		if len(mp) == 1 {
			values[0] = append(values[0], card)
		} else if len(mp) == 2 {
			if mx == 4 {
				values[1-plusTier] = append(values[1-plusTier], card)
			} else {
				idx := int(math.Max(2-float64(plusTier*c), 0))
				values[idx] = append(values[idx], card)
			}
		} else if len(mp) == 3 {
			mx := slices.Max(maps.Values(mp))
			if mx == 3 {
				idx := int(math.Max(3-float64(plusTier*2*c), 0))
				if c == 3 {
					idx = 1
				}
				values[idx] = append(values[idx], card)
			} else {
				idx := int(math.Max(4-float64(plusTier*2*c), 0))
				if c > 1 {
					idx += 1
				}
				values[idx] = append(values[idx], card)
			}
		} else if len(mp) == 4 {
			values[5-plusTier*2] = append(values[5-plusTier*2], card)
		} else {
			values[6-plusTier] = append(values[6-plusTier], card)
		}
	}
	fmt.Println(values)
	r := 1
	slices.Reverse(values[:])
	for _, value := range values {
		if len(value) > 1 {
			for _, d := range compareCards(value) {
				sum += toInt(d[1]) * r
				r += 1
			}
		} else if len(value) == 1 {
			v := toInt(value[0][1])
			sum += v * r
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
		for i, char := range card {
			p := 3 * (4 - i)
			s += int(math.Pow(10, float64(p))) * strengths[string(char)]
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
