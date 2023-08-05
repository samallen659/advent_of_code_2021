package day11

import (
	"github.com/samallen659/advent_of_code/2021/pkg/utils"
	"strconv"
	"strings"
)

var input = utils.ReadInput("/inputs/day11/input1.txt")

func Part1() int {
	grid := buildGrid(input)
	flashCount := 0
	for i := 0; i < 100; i++ {
		var flashedGrid [10][10]bool
		for y := range grid {
			for x := range grid[y] {
				grid[y][x]++
			}
		}
		for y := range grid {
			for x := range grid[y] {
				if grid[y][x] > 9 {
					flash(&grid, y, x, &flashCount, &flashedGrid)
				}
			}
		}
	}
	return flashCount
}

func Part2() int {
	grid := buildGrid(input)
	step := 0
	var flashCount int
	for flashCount < 100 {
		step++
		flashCount = 0
		var flashedGrid [10][10]bool
		for y := range grid {
			for x := range grid[y] {
				grid[y][x]++
			}
		}
		for y := range grid {
			for x := range grid[y] {
				if grid[y][x] > 9 {
					flash(&grid, y, x, &flashCount, &flashedGrid)
				}
			}
		}
	}

	return step
}

func buildGrid(input string) [10][10]int {
	data := strings.Split(input, "\n")
	var grid [10][10]int
	for i, line := range data {
		chars := strings.Split(line, "")
		for j, char := range chars {
			charInt, _ := strconv.Atoi(char)
			grid[i][j] = charInt
		}
	}
	return grid
}

func flash(grid *[10][10]int, y int, x int, flashCount *int, flashedGrid *[10][10]bool) {
	grid[y][x] = 0
	flashedGrid[y][x] = true
	*flashCount++
	for _, yy := range [3]int{-1, 0, 1} {
		for _, xx := range [3]int{-1, 0, 1} {
			if yy != 0 || xx != 0 {
				y1 := y + yy
				x1 := x + xx
				if x1 >= 0 && x1 <= 9 && y1 >= 0 && y1 <= 9 {
					if !flashedGrid[y1][x1] {
						grid[y1][x1]++
						if grid[y1][x1] > 9 {
							flash(grid, y1, x1, flashCount, flashedGrid)
						}
					}
				}
			}
		}
	}
}
