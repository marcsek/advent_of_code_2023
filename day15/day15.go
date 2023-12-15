package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type record struct {
	label string
	fl    string
}

func main() {
	input := readFile("test.txt")
	records := strings.Split(strings.ReplaceAll(input, "\n", ","), ",")
	boxes := map[int][]record{}

	for _, rec := range records[:len(records)-1] {
		addtion := strings.Index(rec, "=") != -1
		label := ""

		if addtion {
			sp := strings.Split(rec, "=")
			nRec := record{sp[0], sp[1]}
			label = sp[0]

			box, ok := boxes[getHash(nRec.label)]
			if ok {
				didFind := false
				for i, r := range box {
					if r.label == nRec.label {
						box[i] = nRec
						didFind = true
						break
					}
				}
				if !didFind {
					box = append(box, nRec)
				}
			} else {
				box = []record{nRec}
			}
			boxes[getHash(label)] = box
		} else {
			label = rec[:len(rec)-1]
			box, ok := boxes[getHash(label)]

			if ok && len(box) > 0 {
				for i, r := range box {
					if label == r.label {
						endIdx := int(math.Min(float64(i+1), float64(len(box))))
						box = append(box[:i], box[endIdx:]...)
						break
					}
				}
			}
			boxes[getHash(label)] = box
		}
	}

	total := 0
	for bNum, bxs := range boxes {
		for i, rec := range bxs {
			total += (bNum + 1) * (i + 1) * toInt(rec.fl)
		}
	}
	fmt.Println(total)
}

func getHash(str string) int {
	curVal := 0

	for _, ch := range str {
		curVal += int(ch)
		curVal *= 17
		curVal %= 256
	}

	return curVal
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
