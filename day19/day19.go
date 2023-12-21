package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type condition struct {
	pr    string
	cond  string
	bound int
	res   string
}

type pathCond struct {
	cond condition
	non  bool
}

var partTypeMap = map[string]int{"x": 0, "m": 1, "a": 2, "s": 3}

func main() {
	input := strings.Split(readFile("test.txt"), "\n")
	workflows := parseInput(input[:len(input)-1])

	total := 0
	path := new([][]pathCond)
	findAllPaths("in", workflows, []pathCond{}, path)
	for _, pth := range *path {
		if pth[len(pth)-1].cond.res != "R" {
			total += evalueatePasses(pth)
		}
	}

	fmt.Println(total)
}

func evalueatePasses(outcomes []pathCond) int {
	values := [4][2]int{{1, 4000}, {1, 4000}, {1, 4000}, {1, 4000}}
	total := 1

	for _, outc := range outcomes {
		if outc.cond.bound == -1 {
			continue
		}
		if !outc.non {
			if outc.cond.cond == "<" {
				target := outc.cond.bound
				cur := &values[partTypeMap[outc.cond.pr]]
				cur[1] = int(math.Min(float64(target), float64(cur[1]))) - 1
			} else if outc.cond.cond == ">" {
				target := outc.cond.bound
				cur := &values[partTypeMap[outc.cond.pr]]
				cur[0] = int(math.Max(float64(target), float64(cur[0]))) + 1
			}
		} else {
			if outc.cond.cond == "<" {
				target := outc.cond.bound
				cur := &values[partTypeMap[outc.cond.pr]]
				cur[0] = int(math.Max(float64(target), float64(cur[0])))
			} else if outc.cond.cond == ">" {
				target := outc.cond.bound
				cur := &values[partTypeMap[outc.cond.pr]]
				cur[1] = int(math.Min(float64(target), float64(cur[1])))
			}
		}
	}
	for _, in := range values {
		total *= (in[1] - in[0]) + 1
	}
	return total
}

func findAllPaths(
	wrkflvId string,
	workFlows map[string][]condition,
	path []pathCond,
	allP *[][]pathCond,
) {
	curwrkflw := workFlows[wrkflvId]

	for i, cond := range curwrkflw {
		curPath := []pathCond{}
		curPath = append(curPath, path...)
		for _, flw := range curwrkflw[:i] {
			curPath = append(curPath, pathCond{flw, true})
		}
		curPath = append(curPath, pathCond{cond, false})
		if cond.res == "A" || cond.res == "R" {
			*allP = append(*allP, curPath)
		} else {
			findAllPaths(cond.res, workFlows, curPath, allP)
		}
	}
}

func parseCond(in string) condition {
	sp := strings.Split(in, ":")
	if len(sp) == 1 {
		return condition{"", "", -1, in}
	}
	pr := string(sp[0][0])
	cond := string(sp[0][1])
	bound := toInt(sp[0][2:])
	res := sp[1]
	return condition{pr, cond, bound, res}
}

func parseInput(input []string) map[string][]condition {
	workflows := map[string][]condition{}

	for _, line := range input {
		if line == "" {
			break
		}
		sp := strings.Split(line, "{")
		id := sp[0]
		wrkfls := []condition{}
		for _, wrkfl := range strings.Split(sp[1][:len(sp[1])-1], ",") {
			wrkfls = append(wrkfls, parseCond(wrkfl))
		}
		workflows[id] = wrkfls
	}
	return workflows
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
