package day12

import (
	"fmt"
	"github.com/samallen659/advent_of_code/2021/pkg/utils"
	"golang.org/x/exp/slices"
	"strings"
	"unicode"
)

var input = utils.ReadInput("/inputs/day12/testInput.txt")
var data = strings.Split(input, "\n")

func Part1() int {
	graph := buildGraph(data)
	fmt.Println(graph)
	var stack Stack
	var visited []string
	pathCount := 0
	stack.push("start")
	for stack.len > 0 {
		cave, err := stack.pop()
		if err != nil {
			break
		}
		if cave == "end" {
			pathCount++
			continue
		}
		visited = append(visited, cave)
		for _, connected := range graph[cave] {
			if connected != "start" {
				if unicode.IsLower([]rune(connected)[0]) && !slices.Contains(visited, connected) {
					stack.push(connected)
				} else {
					stack.push(connected)
				}
			}
		}

		fmt.Println(visited)
	}

	return pathCount
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

type Stack struct {
	data []string
	len  int
}

func (s *Stack) push(d string) {
	s.data = append(s.data, d)
	s.len++
}

func (s *Stack) pop() (string, error) {
	if s.len == 0 {
		return "", fmt.Errorf("Empty Stack")
	}

	val := s.data[s.len-1]
	s.data = append(s.data[:s.len-1], s.data[s.len:]...)
	s.len--

	return val, nil
}
