package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type vec struct {
	y int
	x int
}

var moves = []vec{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
var startPoint = vec{0, 0}
var goal vec

func main() {
	input := strings.Split(readFile("sample.txt"), "\n")
	maze := make([][]int, len(input)-1)
	for y, r := range input[:len(input)-1] {
		maze[y] = append(maze[y], strListToInt(r)...)
	}
	goal = vec{len(maze) - 1, len(maze[len(maze)-1]) - 1}
	goal = vec{1, 3}

	moveMap := walkMazePriority(maze)
	path := reconstructPath(moveMap)
	printWithMove(maze, path)
	fmt.Printf("total cost: %v\n", countMovementCost(maze, path))
}

func walkMazePriority(maze [][]int) map[vec]vec {
	queue := make(PriorityQueue, 0)
	heap.Push(&queue, &Item{value: startPoint, priority: 0})
	cameFrom := map[vec]vec{}
	costSoFar := map[vec]int{}
	cameFrom[startPoint] = vec{}
	costSoFar[startPoint] = 0

	for len(queue) > 0 {
		cur := heap.Pop(&queue).(*Item).value
		if cur.y == goal.y && cur.x == goal.x {
			break
		}
		//fmt.Println(cur, costSoFar[cur], cameFrom)
		for _, m := range moves {
			nY, nX := m.y+cur.y, m.x+cur.x
			if 0 <= nY && len(maze) > nY && 0 <= nX && len(maze[nY]) > nX {
				next := vec{nY, nX}
				newCost := costSoFar[cur] + maze[nY][nX]
				pc := cameFrom[next]
				cameFrom[next] = cur
				if hasStraightLines(maze, cameFrom, next, cur) == 4 {
					continue
				}
				cameFrom[next] = pc
				if cost, ok := costSoFar[next]; !ok || newCost < cost {
					costSoFar[next] = newCost
					heap.Push(&queue, &Item{value: next, priority: newCost})
					cameFrom[next] = cur
				} else if newCost == cost {
					pc := cameFrom[next]
					cameFrom[next] = cur
					t1 := hasStraightLines(maze, cameFrom, next, cur)
					cameFrom[next] = pc
					t2 := hasStraightLines(maze, cameFrom, cameFrom[pc], pc)
					if t2 >= t1 {
						cameFrom[next] = pc
					}
					fmt.Println(t1, t2)
				}
			}
		}
	}
	return cameFrom
}

func hasStraightLines(maze [][]int, moveMap map[vec]vec, start vec, pStart vec) int {
	cur := start
	prev := pStart
	prevDir := vec{cur.y - prev.y, cur.x - prev.x}
	total := 0
	buff := 4

	for (cur.y != startPoint.y || cur.x != startPoint.x) && buff > 0 {
		if cur.x-prev.x == prevDir.x && cur.y-prev.y == prevDir.y {
			total += 1
		} else {
			total = 0
		}
		if total == 4 {
			return 4
		}
		prevDir = vec{cur.y - prev.y, cur.x - prev.x}
		cur = prev
		prev = moveMap[cur]
		buff--
	}
	return total
}

func walkMaze(maze [][]int) map[vec]vec {
	queue := []vec{startPoint}
	cameFrom := map[vec]vec{}
	cameFrom[startPoint] = vec{}

	buffer := -1
	for len(queue) > 0 {
		cur := queue[0]
		if cur.y == goal.y && cur.x == goal.x || buffer == 0 {
			break
		}
		for _, m := range moves {
			nY, nX := m.y+cur.y, m.x+cur.x
			if 0 <= nY && len(maze) > nY && 0 <= nX && len(maze[nY]) > nX {
				next := vec{nY, nX}
				if _, ok := cameFrom[next]; !ok {
					queue = append(queue, vec{nY, nX})
					cameFrom[next] = cur
				}
			}
		}
		buffer--
		queue = queue[1:]
	}
	return cameFrom
}

func reconstructPath(moveMap map[vec]vec) []vec {
	cur := goal
	path := []vec{}

	for cur.y != startPoint.y || cur.x != startPoint.x {
		path = append(path, cur)
		cur = moveMap[cur]
	}
	return path
}

func countMovementCost(maze [][]int, path []vec) int {
	total := 0
	for y, r := range maze {
		for x, e := range r {
			isInPath := false
			for _, p := range path[1:] {
				if p.y == y && p.x == x {
					isInPath = true
					break
				}
			}
			if isInPath {
				total += e
			}
		}
	}
	return total
}

func printWithMove[T any](maze [][]T, path []vec) {
	for y, r := range maze {
		for x, e := range r {
			isInPath := false
			for _, p := range path {
				if p.y == y && p.x == x {
					isInPath = true
					break
				}
			}
			if isInPath {
				fmt.Printf("%v", "#")
			} else {
				fmt.Printf("%v", e)
			}
		}
		fmt.Println()
	}
}

func printMaze[T any](maze [][]T) {
	for _, r := range maze {
		for _, e := range r {
			fmt.Printf(" %2v ", e)
		}
		fmt.Println()
	}
}

func strListToInt(list string) []int {
	c_list := []int{}
	for i := 0; i < len(list); i++ {
		num, err := strconv.ParseInt(string(list[i]), 10, 0)
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

type Item struct {
	value    vec
	priority int
}

type PriorityQueue []*Item

func (piq PriorityQueue) Len() int {
	return len(piq)
}
func (piq PriorityQueue) Less(i, j int) bool {
	return piq[i].priority < piq[j].priority
}
func (piq PriorityQueue) Swap(i, j int) {
	piq[i], piq[j] = piq[j], piq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (piq *PriorityQueue) Pop() interface{} {
	old := *piq
	n := len(old)
	item := old[n-1]
	*piq = old[0 : n-1]
	return item
}
func (piq *PriorityQueue) Update(item *Item, value vec, priority int) {
	item.value = value
	item.priority = priority
}
