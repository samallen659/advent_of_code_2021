package day5

import (
	// "fmt"
	"strconv"
	"strings"

	"github.com/samallen659/advent_of_code/2021/pkg/utils"
)

type point struct {
	x, y int
}

type line struct {
	p1, p2 point
}

type canvas [h][w]int

var can1, can2 canvas

const (
	w = 1000
	h = 1000
)

var input = utils.ReadInput("/inputs/day5/input1.txt")
var lineData = strings.Split(input, "\n")

func Part1() int {
	var lines []line
	for _, l := range lineData {
		newLine := createLine(l)
		lines = append(lines, newLine)
	}
	for _, line := range lines {
		plot(line, false)
	}
	return lineOverlaps(can1)
}

func Part2() int {
	var lines []line
	for _, l := range lineData {
		newLine := createLine(l)
		lines = append(lines, newLine)
	}
	for _, line := range lines {
		plot(line, true)
	}
	return lineOverlaps(can2)
}

func lineOverlaps(can canvas) int {
	overLapCount := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if can[i][j] > 1 {
				overLapCount++
			}
		}
	}

	return overLapCount
}

func plot(l line, diag bool) {
	deltaX, deltaY := l.p1.x-l.p2.x, l.p1.y-l.p2.y

	switch {
	case deltaX == 0:
		minY, maxY := sort(l.p1.y, l.p2.y)
		for i := minY; i < maxY+1; i++ {
			can1[i][l.p1.x]++
			if diag {
				can2[i][l.p1.x]++
			}
		}
	case deltaY == 0:
		minX, maxX := sort(l.p1.x, l.p2.x)
		for i := minX; i < maxX+1; i++ {
			can1[l.p1.y][i]++
			if diag {
				can2[l.p1.y][i]++
			}
		}
	case diag:
		//slope intercept
		m := deltaY / deltaX
		b := l.p1.y - l.p1.x*m
		minX, maxX := sort(l.p1.x, l.p2.x)
		for x := minX; x < maxX+1; x++ {
			y := m*x + b
			can2[y][x]++
		}
	}

}

func createLine(lineStr string) line {
	pointData := strings.Split(lineStr, " -> ")
	var newLine line
	for i, p := range pointData {
		points := strings.Split(p, ",")
		x, _ := strconv.Atoi(points[0])
		y, _ := strconv.Atoi(points[1])
		if i == 0 {
			newLine.p1.x = x
			newLine.p1.y = y
		} else {
			newLine.p2.x = x
			newLine.p2.y = y
		}
	}
	return newLine
}

func sort(n1, n2 int) (int, int) {
	if n1 < n2 {
		return n1, n2
	}
	return n2, n1
}
