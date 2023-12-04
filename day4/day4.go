package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := readFile("test.txt")
	cards := strings.Split(input, "\n")
	cards = cards[:len(cards)-1]
	queue := []int{}
	lenCards := len(cards)

	for range cards {
		queue = append(queue, 1)
	}
	in_queue := len(queue)

	for in_queue > 0 {
		for q_id, q_amount := range queue {
			card := cards[q_id]
			for i := 0; i < q_amount; i++ {
				nums := strings.Split(card, ": ")[1]
				wnums := strings.Split(nums, " | ")[0]
				mnums := strings.Split(nums, " | ")[1]
				wins := 0
				in_queue--

				for _, wnum := range strings.Split(wnums, " ") {
					for _, mnum := range strings.Split(mnums, " ") {
						if wnum == mnum && wnum != "" {
							wins += 1
						}
					}
				}
				lenCards += wins
				for q := q_id + 1; q <= q_id+wins; q++ {
					queue[q]++
					in_queue++
				}
			}
		}
	}
	fmt.Println(lenCards)
}

func readFile(fn string) string {
	data, err := os.ReadFile(fn)

	if err != nil {
		panic("Couldn't read input file")
	}

	return string(data)
}
