package day12

import (
	"fmt"
	"strings"

	"github.com/samallen659/advent_of_code/2021/pkg/utils"
	"golang.org/x/exp/slices"
)

var input = utils.ReadInput("/inputs/day12/testInput.txt")
var data = strings.Split(input, "\n")

func Part1() {
	graph := buildGraph(data)
	var stack Stack
	var visited []string
	stack.push("start")
	for stack.len > 0 {
		cave, err := stack.pop()
		visited = append(visited, cave)
		if err != nil {
			break
		}
		for _, connected := range graph[cave] {
			fmt.Println(connected)
		}

		fmt.Println(cave)
	}
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
		if !slices.Contains(graph[c2], c2) {
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
