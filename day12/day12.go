package main

import (
	"fmt"
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

//???.###? ???.####
//.??..??...?##. ?.??..??...?##.

func main() {
	input := strings.Split(readFile("test.txt"), "\n")
	groups := []group{}

	for _, line := range input[:len(input)-1] {
		sp := strings.Split(line, " ")

		recNN := []rune{}
		recN := []rune{}
		if sp[0][len(sp[0])-1] != '#' {
			recN = append(recN, '?')
		}
		//recN := []rune{'?'}
		recN = append(recN, []rune(sp[0])...)
		recNN = append(recNN, []rune(sp[0])...)
		//recN = append(recN, '?')
		if sp[0][len(sp[0])-1] != '#' {
			if sp[0][len(sp[0])-1] != '?' {
				rc := recNN
				recNN = append([]rune{'?'}, rc...)
			}
			recN = getFiveMore(recN, 1)
		} else {
			recNN = append(recNN, '?')
			recN = getFiveMore(recN, 1, '?')
		}
		recNNN := []rune{'?'}
		if sp[0][len(sp[0])-1] == '#' {
			recNNN = []rune{}
		}
		recNNN = append(recNNN, []rune(sp[0])...)
		rec := getFiveMore([]rune(sp[0]), 1, '?')

		posN := strListToInt(strings.Split(sp[1], ","))
		if posN[len(posN)-1] == 1 && sp[0][len(sp[0])-1] == '?' {
			recNNN = recNNN[1:]
			recNNN = append(recNNN, '?')
			rec = append(rec, '?')
			recNN = append(recNN, '?')
		}
		pos := getFiveMore(posN, 1)
		posN = getFiveMore(posN, 1)
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
		quesm = []int{}
		for x, r := range recNN {
			if r == '?' {
				quesm = append(quesm, x)
			}
		}
		groups = append(groups, group{Rec: []rune(recNN), Pos: posN, Quesm: quesm})

		quesm = []int{}
		for x, r := range recNNN {
			if r == '?' {
				quesm = append(quesm, x)
			}
		}
		groups = append(groups, group{Rec: []rune(recNNN), Pos: pos, Quesm: quesm})
	}

	total := 0
	for i := 0; i < len(groups); i += 4 {
		d := genValid(groups[i+2])
		p1 := genValid(groups[i+1])
		p2 := genValid(groups[i])
		p3 := d * d
		p4 := genValid(groups[i+3])
		//fmt.Println(p1 * p2 * p3 * p4)
		total += p1 * p2 * p3 * p4
	}
	//fmt.Println(string(groups[12].Rec))
	//fmt.Println(genValid(groups[12]))
	//fmt.Println(string(groups[13].Rec))
	//fmt.Println(genValid(groups[13]))
	//fmt.Println(string(groups[14].Rec))
	//fmt.Println(genValid(groups[14]))
	//fmt.Println(string(groups[15].Rec))
	//fmt.Println(genValid(groups[15]))
	fmt.Println(total)
}

func getFiveMore[T any](sl []T, m int, z ...T) []T {
	newSl := []T{}

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
