package day15

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samallen659/advent_of_code/2021/pkg/utils"
)

var input = utils.ReadInput("/inputs/day15/testInput.txt")
var data = strings.Split(input, "\n")

func Part1() {
	graph := buildGraph(data)
}

func buildGraph(data []string) *[][]int {
	n := len(data)
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, len(data[0]))
	}
	for i, line := range data {
		chars := strings.Split(line, "")
		for j, char := range chars {
			val, _ := strconv.Atoi(char)
			graph[i][j] = val
		}
	}

	return &graph
}
