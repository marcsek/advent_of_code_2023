package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := strings.Split(readFile("test.txt"), "\n")
	seques := make([][][]int, len(input)-1)

	for i, line := range input[:len(input)-1] {
		seques[i] = append(seques[i], strListToInt(strings.Split(line, " ")))
	}
	s := 0
	for _, seq := range seques {
		v := genSeq(seq)
		s += v
	}

	fmt.Println(s)
}

func genSeq(seq [][]int) int {
	sIdx := 0
	allZero := false
	for !allZero {
		nSeq := []int{}
		allZero = true
		for i := 0; i < len(seq[sIdx])-1; i++ {
			p0 := seq[sIdx][i]
			p1 := seq[sIdx][i+1]
			nSeq = append(nSeq, p1-p0)
			if p1-p0 != 0 {
				allZero = false
			}
		}
		seq = append(seq, nSeq)
		sIdx++
	}

	lastV := 0
	slices.Reverse(seq)
	for _, set := range seq {
		lastV = set[0] - lastV
	}

	return lastV
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
		num, err := strconv.ParseInt(list[i], 10, 0)
		if err != nil {
			panic("Couldn't parse int")
		}
		c_list = append(c_list, int(num))
	}
	return c_list
}
