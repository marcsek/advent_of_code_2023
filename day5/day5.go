package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := readFile("test.txt")

	lines := strings.Split(input, "\n")
	seeds := mkSeedRng(strListToInt(strings.Split(lines[0], " ")[1:]))
	rangs := parseRanges(lines)
	locations := []int{}

	for _, seed := range seeds {
		seed_v := seed
		for _, rang := range rangs {
			for _, mp := range rang {
				drng := mp[0]
				srng := mp[1]
				rng := mp[2]

				if srng <= seed_v && seed_v <= srng+rng {
					seed_v = drng + (seed_v - srng)
					break
				}
			}
		}
		locations = append(locations, seed_v)
	}
	fmt.Println(slices.Min(locations))
}

func mkSeedRng(seeds []int) []int {
	rngSeeds := []int{}
	for i := 0; i < len(seeds); i += 2 {
		fmt.Print(".")
		for j := seeds[i]; j < seeds[i+1]+seeds[i]; j++ {
			rngSeeds = append(rngSeeds, j)
		}
	}
	return rngSeeds
}

func parseRanges(lines []string) [][][]int {
	rangs := [][][]int{}
	rang := [][]int{}
	je_nadpis := true

	for _, line := range lines[2:] {
		if je_nadpis {
			je_nadpis = false
			continue
		}
		if line == "" {
			je_nadpis = true
			rangs = append(rangs, rang)
			rang = [][]int{}
			continue
		}
		rang = append(rang, strListToInt(strings.Split(line, " ")))
	}
	return rangs
}

func strListToInt(list []string) []int {
	c_list := []int{}
	for i := 0; i < len(list); i++ {
		num, err := strconv.ParseInt(list[i], 10, 0)
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
