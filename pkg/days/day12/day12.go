package day12

import (
	"github.com/samallen659/advent_of_code/2021/pkg/utils"
	"golang.org/x/exp/slices"
	"strings"
	"unicode"
)

var input = utils.ReadInput("/inputs/day12/input1.txt")
var data = strings.Split(input, "\n")
var graph = buildGraph(data)

func Part1() int {
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

func Part2() int {
	pathCount := 0
	visited := make(map[string]int)
	var dfs func(string)
	dfs = (func(node string) {
		if node == "end" {
			pathCount++
			return
		}
		if isLower(node) {
			visited[node] += 1
		}
		for _, neighbor := range graph[node] {
			if neighbor != "start" {
				if visitedSmallCaveTwice(&visited) {
					if visited[neighbor] == 0 {
						dfs(neighbor)
					}
				} else {
					dfs(neighbor)
				}

			}
		}
		if isLower(node) {
			visited[node] -= 1
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

func visitedSmallCaveTwice(visited *map[string]int) bool {
	for _, v := range *visited {
		if v == 2 {
			return true
		}
	}
	return false
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
