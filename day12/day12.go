package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

type group struct {
	Rec   []rune
	Pos   []int
	Quesm []int
}

func main() {
	input := strings.Split(readFile("sample.txt"), "\n")
	groups := []group{}

	for _, line := range input[:len(input)-1] {
		sp := strings.Split(line, " ")
		recN := []rune(sp[0])
		recN = append(recN, '?')
		rec := getFiveMore([]rune(sp[0]), '?')
		posN := strListToInt(strings.Split(sp[1], ","))
		pos := getFiveMore(posN)
		quesm := []int{}
		for x, r := range rec {
			if r == '?' {
				quesm = append(quesm, x)
			}
		}
		groups = append(groups, group{Rec: []rune(rec), Pos: pos, Quesm: quesm})
		quesm = []int{}
		for x, r := range recN {
			if r == '?' {
				quesm = append(quesm, x)
			}
		}
		groups = append(groups, group{Rec: []rune(recN), Pos: posN, Quesm: quesm})
	}

	total := 0
	for i := 0; i < len(groups); i += 2 {
		p1 := int(math.Pow(float64(genValid(groups[i+1])), 3))
		p2 := genValid(groups[i])
		fmt.Println(p1, p2)
		total += p1 * p2
	}
	fmt.Println(total)
}

func getFiveMore[T any](sl []T, z ...T) []T {
	newSl := []T{}
	m := 2

	for i := 0; i < m; i++ {
		newSl = append(newSl, sl...)
		if len(z) != 0 && i < m-1 {
			newSl = append(newSl, z...)
		}
	}
	return newSl
}

func genValid(g group) int {
	left := sliceSum(g.Pos) - genAmount(g)
	comb := combin.Combinations(len(g.Quesm), left)
	c := 0

	for _, p := range comb {
		newStr := slices.Clone(g.Rec)
		for _, x := range p {
			cor := g.Quesm[x]
			newStr[cor] = '#'
		}
		if checkIfValid(newStr, g.Pos) {
			c++
		}
	}
	return c
}

func sliceSum(s []int) int {
	total := 0
	for _, e := range s {
		total += e
	}
	return total
}

func genAmount(g group) int {
	am := 0
	for _, r := range g.Rec {
		if r == '#' {
			am++
		}
	}
	return am
}

func checkIfValid(rec []rune, pos []int) bool {
	cur := 0
	i := 0

	for _, r := range rec {
		if r == '#' {
			cur++
		} else if cur != 0 {
			if cur != pos[i] {
				return false
			}
			cur = 0
			i++
		}
	}
	if cur != 0 && cur != pos[len(pos)-1] {
		return false
	}
	return true
}

func strListToInt(list []string) []int {
	c_list := []int{}
	for i := 0; i < len(list); i++ {
		if list[i] == "" {
			continue
		}
		num, err := strconv.ParseInt(strings.TrimSpace(list[i]), 10, 0)
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
