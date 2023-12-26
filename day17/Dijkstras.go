package main

import (
	"container/heap"
)

type Vec struct {
	y int
	x int
}

type baseState comparable

type Dijkstras[T any] interface {
	neighbours(cur T) []T
	goalReached(cur T, goal T) bool
	costToMove(from T, to T) int
	distance(from T, to T) int
	isInside(point Vec) bool
}

type pathNode[T any] struct {
	self      T
	parent    *pathNode[T]
	costSoFar int
}

func shortestPath[T baseState](maze Dijkstras[T], from T, dest T) []T {
	frontier := make(PriorityQueue[T], 0)
	heap.Push(&frontier, &Item[T]{value: pathNode[T]{self: from}, priority: 0})
	explored := map[T]int{from: 0}

	for len(frontier) > 0 {
		curNode := &heap.Pop(&frontier).(*Item[T]).value
		cur := curNode.self

		if maze.goalReached(cur, dest) {
			res := []T{}

			node := curNode
			for node.parent != nil {
				res = append(res, node.self)
				node = node.parent
			}
			return res[1:]
		}

		for _, n := range maze.neighbours(cur) {
			moveCost := maze.costToMove(cur, n)
			newCost := curNode.costSoFar + moveCost

			if cost, ok := explored[n]; !ok || cost > newCost {
				explored[n] = newCost
				heap.Push(&frontier, &Item[T]{value: pathNode[T]{self: n, parent: curNode, costSoFar: newCost}, priority: newCost})
			}
		}
	}

	return []T{from, dest}
}

type Item[T baseState] struct {
	value    pathNode[T]
	priority int
}

type PriorityQueue[T baseState] []*Item[T]

func (piq PriorityQueue[T]) Len() int {
	return len(piq)
}

func (piq PriorityQueue[T]) Less(i, j int) bool {
	return piq[i].priority < piq[j].priority
}

func (piq PriorityQueue[T]) Swap(i, j int) {
	piq[i], piq[j] = piq[j], piq[i]
}

func (pq *PriorityQueue[T]) Push(x interface{}) {
	*pq = append(*pq, x.(*Item[T]))
}

func (piq *PriorityQueue[T]) Pop() interface{} {
	old := *piq
	n := len(old)
	item := old[n-1]
	*piq = old[0 : n-1]
	return item
}
