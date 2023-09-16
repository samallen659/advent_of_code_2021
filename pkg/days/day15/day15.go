package day15

import (
	"container/heap"
	"math"
	"strconv"
	"strings"

	"github.com/samallen659/advent_of_code/2021/pkg/utils"
)

var input = utils.ReadInput("/inputs/day15/input1.txt")
var data = strings.Split(input, "\n")

func Part1() int {
	graph := buildGraph(data)
	risk := getRisk(graph)
	return risk
}

func Part2() int {
	graph := buildGraph2(data)
	risk := getRisk(graph)
	return risk
}

func buildGraph(data []string) *[][]int {
	n, m := len(data), len(data[0])
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, m)
		line := strings.Split(data[i], "")
		for j, char := range line {
			val, _ := strconv.Atoi(char)
			graph[i][j] = val
		}
	}
	return &graph
}

func buildGraph2(data []string) *[][]int {
	n, m := len(data), len(data[1])
	graph := make([][]int, (n * 5))
	iAddr := 0
	for i := 0; i < (n * 5); i++ {
		graph[i] = make([]int, (m * 5))
		line := strings.Split(data[i%n], "")
		if i > 0 && i%n == 0 {
			iAddr++
		}
		jAddr := 0
		for j := 0; j < (m * 5); j++ {
			if j >= m && j%m == 0 {
				jAddr++
			}
			val, _ := strconv.Atoi(line[j%m])
			val = val + iAddr + jAddr
			if val > 9 {
				val = 0 + (val - 9)
			}
			graph[i][j] = val
		}
	}
	return &graph
}

// dijkstra
func getRisk(graph *[][]int) int {
	n, m := len(*graph), len((*graph)[0])

	dist := make(map[Node]int)
	pq := make(PriorityQueue, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			tempNode := Node{y: i, x: j}
			dist[tempNode] = math.MaxInt32
			pq[i*m+j] = (&Item{
				value:    tempNode,
				priority: math.MaxInt32,
				index:    i*m + j,
			})
		}
	}

	heap.Init(&pq)
	src := Node{y: 0, x: 0}
	dist[src] = 0
	pq.update(src, 0)
	neighborNodes := []Node{{y: 0, x: 1}, {y: 0, x: -1}, {y: 1, x: 0}, {y: -1, x: 0}}

	for len(pq) > 0 {
		item := pq.Pop().(*Item)

		for _, neighbor := range neighborNodes {
			node := Node{y: neighbor.y + item.value.y, x: neighbor.x + item.value.x}
			if node.y >= 0 && node.y < n && node.x >= 0 && node.x < m {
				if dist[item.value]+(*graph)[node.y][node.x] <
					dist[node] {
					dist[node] = dist[item.value] + (*graph)[node.y][node.x]
					pq.update(node, dist[node])
				}
			}
		}
	}

	return dist[Node{y: n - 1, x: m - 1}]
}
