package day15

import (
	"container/heap"
	"fmt"
)

type Node struct {
	y int
	x int
}

type Item struct {
	value    Node
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	pq.Swap(0, n-1)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	heap.Fix(pq, 0)
	return item
}

func (pq *PriorityQueue) updateByItem(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

func (pq *PriorityQueue) update(value Node, priority int) {
	for _, qi := range *pq {
		if qi.value == value {
			pq.updateByItem(qi, priority)
			return
		}
	}
}

func (pq *PriorityQueue) Print() {
	for _, qi := range *pq {
		fmt.Println(qi)
	}
}
