package day12

import (
	"github.com/samallen659/advent_of_code/2021/pkg/utils"
	"golang.org/x/exp/slices"
	"strings"
	"unicode"
)

var input = utils.ReadInput("/inputs/day12/input1.txt")
var data = strings.Split(input, "\n")

func Part1() int {
	graph := buildGraph(data)
	pathCount := 0
	visited := make(map[string]bool)
	var dfs func(string)
	dfs = (func(node string) {
		if node == "end" {
			pathCount++
			return
		}
		if isLower(node) {
			visited[node] = true
		}
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
		if isLower(node) {
			visited[node] = false
		}
	})

	dfs("start")

	return pathCount
}

func isLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func buildGraph(input []string) map[string][]string {
	graph := make(map[string][]string)

	for _, line := range data {
		caves := strings.Split(line, "-")
		c1 := caves[0]
		c2 := caves[1]
		if !slices.Contains(graph[c1], c2) {
			graph[c1] = append(graph[c1], c2)
		}
		if !slices.Contains(graph[c2], c1) {
			graph[c2] = append(graph[c2], c1)
		}
	}
	return graph
}
