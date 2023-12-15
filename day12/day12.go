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

// .??..??...?##.
// .??..??...?##.?.??..??...?##. [1]

// ?#?#?#?#?#?#?#?
// ?#?#?#?#?#?#?#???#?#?#?#?#?#?#? [1]

// ?????.#...#...

type group struct {
	rec   []rune
	pos   []int
	quesm []int
}

func main() {
	input := strings.Split(readfile("test.txt"), "\n")
	groups := []group{}

	for _, line := range input[:len(input)-1] {
		sp := strings.Split(line, " ")

		rec := []rune{}
		rec = append(rec, []rune(sp[0])...)
		pos := strlisttoint(strings.Split(sp[1], ","))

		prepMod := []rune{'?'}
		appMod := []rune(sp[0])
		appMod = append(appMod, '?')
		prepMod = append(prepMod, []rune(sp[0])...)

		quesm := []int{}
		for x, r := range rec {
			if r == '?' {
				quesm = append(quesm, x)
			}
		}
		groups = append(groups, group{rec: []rune(rec), pos: pos, quesm: quesm})

		quesm = []int{}
		for x, r := range prepMod {
			if r == '?' {
				quesm = append(quesm, x)
			}
		}
		groups = append(groups, group{rec: []rune(prepMod), pos: pos, quesm: quesm})

		quesm = []int{}
		for x, r := range appMod {
			if r == '?' {
				quesm = append(quesm, x)
			}
		}
		groups = append(groups, group{rec: []rune(appMod), pos: pos, quesm: quesm})
	}

	total := 0
	for i := 0; i < len(groups); i += 3 {
		orgGroup := groups[i]
		prepGroup := groups[i+1]
		appGroup := groups[i+2]

		// asi treba zdokonalit podmienku komplementu
		validCompl := genvalid(appGroup)
		canNegate := orgGroup.rec[len(orgGroup.rec)-1] == '#' || orgGroup.rec[0] == '#'

		if potCompl := genvalid(prepGroup); potCompl > validCompl && !canNegate {
			validCompl = potCompl
		}
		//fmt.Println(validCompl, orgGroup)

		//fmt.Println(genvalid(orgGroup) * int(math.Pow(float64(genvalid(modGroup)), 4)))
		//fmt.Println(
		//	string(orgGroup.rec),
		//	genvalid(orgGroup),
		//	string(modGroup.rec),
		//	genvalid(modGroup),
		//	orgGroup.pos,
		//	isChangable(string(orgGroup.rec), orgGroup.pos[0]),
		//)

		total += genvalid(orgGroup) * int(math.Pow(float64(validCompl), 4))
	}
	fmt.Println(total)
}

func isChangable(rec string, num int) bool {
	total := 1
	if rec[len(rec)-1] == '#' {
		return false
	}

	for _, r := range rec {
		if r != '?' {
			if r == '#' {
				total--
			}
			break
		}
		total++
	}

	return num <= total
}

func getfivemore[t any](sl []t, m int, z ...t) []t {
	newsl := []t{}

	for i := 0; i < m; i++ {
		newsl = append(newsl, sl...)
		if len(z) != 0 && i < m-1 {
			newsl = append(newsl, z...)
		}
	}
	return newsl
}

func genvalid(g group) int {
	left := slicesum(g.pos) - genamount(g)
	comb := combin.Combinations(len(g.quesm), left)
	c := 0

	for _, p := range comb {
		newstr := slices.Clone(g.rec)
		for _, x := range p {
			cor := g.quesm[x]
			newstr[cor] = '#'
		}
		if checkifvalid(newstr, g.pos) {
			c++
		}
	}
	return c
}

func slicesum(s []int) int {
	total := 0
	for _, e := range s {
		total += e
	}
	return total
}

func genamount(g group) int {
	am := 0
	for _, r := range g.rec {
		if r == '#' {
			am++
		}
	}
	return am
}

func checkifvalid(rec []rune, pos []int) bool {
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

func strlisttoint(list []string) []int {
	c_list := []int{}
	for i := 0; i < len(list); i++ {
		if list[i] == "" {
			continue
		}
		num, err := strconv.ParseInt(strings.TrimSpace(list[i]), 10, 0)
		if err != nil {
			panic("couldn't parse int")
		}
		c_list = append(c_list, int(num))
	}
	return c_list
}

func readfile(fn string) string {
	data, err := os.ReadFile(fn)

	if err != nil {
		panic("couldn't read input file")
	}

	return string(data)
}
