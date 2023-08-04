package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samallen659/advent_of_code/2021/pkg/utils"
)

var input = utils.ReadInput("/inputs/day11/testInput.txt")

func Part1() {
	grid := buildGrid(input)
	flashCount := 0
	for i := 0; i < 5; i++ {
		for y := range grid {
			for x := range grid[y] {
				grid[y][x]++
				if grid[y][x] == 10 {
					flash(&grid, y, x, &flashCount)
				}
			}
		}
		// for y := range grid {
		// 	for x := range grid[y] {
		// 	}
		// }
		fmt.Println(grid[0])
	}
	fmt.Println(flashCount)
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

func flash(grid *[10][10]int, y int, x int, flashCount *int) {
	*flashCount++
	grid[y][x] = 0
	for _, yy := range [3]int{-1, 0, 1} {
		for _, xx := range [3]int{-1, 0, 1} {
			if yy != 0 && xx != 0 {
				y1 := y + yy
				x1 := x + xx
				if x1 >= 0 && x1 <= 9 && y1 >= 0 && y1 <= 9 {
					grid[y1][x1]++
					if grid[y1][x1] == 10 {
						flash(grid, y1, x1, flashCount)
					}
				}
			}
		}
	}
}
