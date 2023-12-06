package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := strings.Split(readFile("test.txt"), "\n")
	tms := strListToInt(strings.Split(strings.Replace(input[0][5:], " ", "", -1), " "))
	dsts := strListToInt(strings.Split(strings.Replace(input[1][9:], " ", "", -1), " "))

	total := 1
	for i := 0; i < len(dsts); i++ {
		if race := tryAllRaces(tms[i], dsts[i]); race > 0 {
			total *= tryAllRaces(tms[i], dsts[i])
		}
	}
	fmt.Println(total)
}

func tryAllRaces(time int, record int) int {
	cBst := 0
	for i := 1; i <= time; i++ {
		dst := (time - i) * i
		if dst > record {
			cBst++
		}
	}
	return cBst
}

func readFile(fn string) string {
	data, err := os.ReadFile(fn)

	if err != nil {
		panic("Couldn't read input file")
	}

	return string(data)
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
