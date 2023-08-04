package day1

import (
	"strconv"
	"strings"

	"github.com/samallen659/advent_of_code/2021/pkg/utils"
)

func Part1() int {
	input := utils.ReadInput("/inputs/day1/input1.txt")
	data := strings.Split(input, "\n")
	var increaseCount int
	for i := 1; i < len(data); i++ {
		d1, _ := strconv.Atoi(data[i-1])
		d2, _ := strconv.Atoi(data[i])
		if d1 < d2 {
			increaseCount++
		}
	}
	return increaseCount
}

func Part2() int {
	input := utils.ReadInput("/inputs/day1/input1.txt")
	data := strings.Split(input, "\n")
	var increaseCount int
	for i := 3; i < len(data); i++ {
		d1, _ := strconv.Atoi(data[i-3])
		d2, _ := strconv.Atoi(data[i-2])
		d3, _ := strconv.Atoi(data[i-1])
		d4, _ := strconv.Atoi(data[i])

		sum1 := d1 + d2 + d3
		sum2 := d2 + d3 + d4

		if sum1 < sum2 {
			increaseCount++
		}
	}
	return increaseCount
}
