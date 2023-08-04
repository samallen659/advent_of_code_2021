package day2

import (
	"github.com/samallen659/advent_of_code/2021/pkg/utils"
	"strconv"
	"strings"
)

func Part1() int {
	input := utils.ReadInput("/inputs/day2/input1.txt")
	data := strings.Split(input, "\n")
	horizontalPosition := 0
	depth := 0
	for i := 0; i < len(data); i++ {
		horizontalChange, depthChange := calculateMove(data[i])
		horizontalPosition += horizontalChange
		depth += depthChange
	}
	return depth * horizontalPosition
}

func Part2() int {
	input := utils.ReadInput("/inputs/day2/input1.txt")
	data := strings.Split(input, "\n")
	horizontalPosition := 0
	depth := 0
	aim := 0
	for i := 0; i < len(data); i++ {
		moveData := strings.Split(data[i], " ")
		moveDirection := moveData[0]
		moveAmount, _ := strconv.Atoi(moveData[1])
		switch moveDirection {
		case "forward":
			horizontalPosition += moveAmount
			depth += moveAmount * aim
		case "down":
			aim += moveAmount
		case "up":
			aim -= moveAmount
		}
	}
	return depth * horizontalPosition
}

func calculateMove(move string) (int, int) {
	moveData := strings.Split(move, " ")
	moveDirection := moveData[0]
	moveAmount := moveData[1]

	var horizontalChange int
	var depthChange int

	switch moveDirection {
	case "forward":
		changeInt, _ := strconv.Atoi(moveAmount)
		horizontalChange += changeInt
	case "down":
		changeInt, _ := strconv.Atoi(moveAmount)
		depthChange += changeInt
	case "up":
		changeInt, _ := strconv.Atoi(moveAmount)
		depthChange -= changeInt
	}

	return horizontalChange, depthChange
}
