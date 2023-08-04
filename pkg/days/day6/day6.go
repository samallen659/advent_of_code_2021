package day6

import (
	"github.com/samallen659/advent_of_code/2021/pkg/utils"
	"strconv"
	"strings"
)

var input = utils.ReadInput("/inputs/day6/input1.txt")
var data = strings.Split(input, ",")

func Part1() int {
	var lanternFish []int
	for _, fishStr := range data {
		fishNum, _ := strconv.Atoi(fishStr)
		lanternFish = append(lanternFish, fishNum)
	}

	for i := 0; i < 80; i++ {
		for i, fish := range lanternFish {
			if fish == 0 {
				lanternFish[i] = 6
				lanternFish = append(lanternFish, 8)
			} else {
				lanternFish[i]--
			}
		}
	}
	return len(lanternFish)
}

func Part2() uint64 {
	var lanternFish [9]uint64
	for _, fishStr := range data {
		fishNum, _ := strconv.Atoi(fishStr)
		lanternFish[fishNum]++
	}

	for i := 0; i < 256; i++ {
		updateLanternFish(lanternFish[:])
	}
	var lanternFishCount uint64
	for i := range lanternFish {
		lanternFishCount += lanternFish[i]
	}
	return lanternFishCount
}

func updateLanternFish(l []uint64) {
	i := l[0]
	copy(l, l[1:])
	l[6] += i
	l[8] = i
}
