package day15

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/samallen659/advent_of_code/2021/pkg/utils"
)

var input = utils.ReadInput("/inputs/day15/input1.txt")
var data = strings.Split(input, "\n")

func Part1() {
	graph := buildGraph(data)
	risk := lowestRiskPath(graph)
	fmt.Println(risk)
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

func lowestRiskPath(graph *[][]int) int {
	n, m := len(*graph), len((*graph)[0])

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	dp[0][0] = 0
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + (*graph)[i][0]
	}
	for i := 1; i < m; i++ {
		dp[0][i] = dp[0][i-1] + (*graph)[0][i]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < len(dp[0]); j++ {
			dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))) + (*graph)[i][j]
		}
	}

	return dp[n-1][m-1]
}
