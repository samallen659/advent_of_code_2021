package day7

import (
	// "fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/samallen659/advent_of_code/2021/pkg/utils"
)

var input = utils.ReadInput("/inputs/day7/input1.txt")
var data = strings.Split(input, ",")

func Part1() int {
	var crabPosition []int
	for _, posStr := range data {
		posInt, _ := strconv.Atoi(posStr)
		crabPosition = append(crabPosition, posInt)
	}
	sort.Ints(crabPosition)
	var med = crabPosition[int(math.Round(float64(len(crabPosition)/2)))]

	fuelTotal := 0
	for _, posInt := range crabPosition {
		fuelTotal += Abs(posInt - med)
	}

	return fuelTotal
}

func Part2() int {
	var crabPosition []int
	posTotal := 0
	for _, posStr := range data {
		posInt, _ := strconv.Atoi(posStr)
		posTotal += posInt
		crabPosition = append(crabPosition, posInt)
	}

	sort.Ints(crabPosition)
	minPos, maxPos := crabPosition[0], crabPosition[len(crabPosition)-1]

	minFuelTotal := 2147483647
	for i := minPos; i < maxPos+1; i++ {
		fuelTotal := 0
		for j := 0; j < len(crabPosition); j++ {
			absDist := Abs(i - crabPosition[j])
			fuelUsed := 0
			for f := 0; f < absDist; f++ {
				fuelUsed += f + 1
			}
			fuelTotal += fuelUsed
		}
		if fuelTotal < minFuelTotal {
			minFuelTotal = fuelTotal
		}
	}

	return minFuelTotal
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
