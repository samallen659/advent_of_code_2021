package day9

import (
	"github.com/samallen659/advent_of_code/2021/pkg/utils"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	y int
	x int
}

func createPoint(y int, x int) Point {
	var p Point
	p.y = y
	p.x = x
	return p
}

var input = utils.ReadInput("/inputs/day9/input1.txt")
var lines = strings.Split(input, "\n")
var heightMap = createHeightMap(lines)

func Part1() int {
	lowPoints := findLowPoints(heightMap)
	total := 0
	for _, point := range lowPoints {
		total += point + 1
	}

	return total
}

func Part2() int {
	lowPoints := findLowPointsCoordinates(heightMap)
	var basonSizes []int
	for _, point := range lowPoints {
		bason := calculateBason(point, heightMap)
		basonSizes = append(basonSizes, len(bason))
	}
	sort.Ints(basonSizes)
	result := 1
	for i := len(basonSizes) - 3; i < len(basonSizes); i++ {
		result *= basonSizes[i]
	}
	return result
}

func createHeightMap(lines []string) [][]int {
	var heightMap [][]int
	for _, line := range lines {
		var lineMap []int
		lineSplit := strings.Split(line, "")
		for _, height := range lineSplit {
			heightInt, _ := strconv.Atoi(height)
			lineMap = append(lineMap, heightInt)
		}
		heightMap = append(heightMap, lineMap)
	}
	return heightMap
}

func findLowPoints(heightMap [][]int) []int {
	var lowPoints []int
	for y := 0; y < len(heightMap); y++ {
		for x := 0; x < len(heightMap[y]); x++ {
			switch {
			case y > 0 && heightMap[y-1][x] <= heightMap[y][x]:
				continue
			case y < len(heightMap)-1 && heightMap[y+1][x] <= heightMap[y][x]:
				continue
			case x > 0 && heightMap[y][x-1] <= heightMap[y][x]:
				continue
			case x < len(heightMap[y])-1 && heightMap[y][x+1] <= heightMap[y][x]:
				continue
			default:
				lowPoints = append(lowPoints, heightMap[y][x])
			}
		}
	}
	return lowPoints
}

func findLowPointsCoordinates(heightMap [][]int) []Point {
	var lowPoints []Point
	for y := 0; y < len(heightMap); y++ {
		for x := 0; x < len(heightMap[y]); x++ {
			switch {
			case y > 0 && heightMap[y-1][x] <= heightMap[y][x]:
				continue
			case y < len(heightMap)-1 && heightMap[y+1][x] <= heightMap[y][x]:
				continue
			case x > 0 && heightMap[y][x-1] <= heightMap[y][x]:
				continue
			case x < len(heightMap[y])-1 && heightMap[y][x+1] <= heightMap[y][x]:
				continue
			default:
				lowPoint := createPoint(y, x)
				lowPoints = append(lowPoints, lowPoint)
			}
		}
	}
	return lowPoints
}

func calculateBason(lowPoint Point, heightMap [][]int) []Point {
	visited := make(map[Point]bool)
	var bason []Point
	var queue Queue
	queue.enqueue(lowPoint)
	for queue.len > 0 {
		point := queue.dequeue()
		if !visited[point] {
			bason = append(bason, point)
			visited[point] = true
			abovePoint := createPoint(point.y-1, point.x)
			leftPoint := createPoint(point.y, point.x-1)
			belowPoint := createPoint(point.y+1, point.x)
			rightPoint := createPoint(point.y, point.x+1)
			if abovePoint.y >= 0 && pointInBason(point, abovePoint, heightMap) {
				queue.enqueue(abovePoint)
			}
			if leftPoint.x >= 0 && pointInBason(point, leftPoint, heightMap) {
				queue.enqueue(leftPoint)
			}
			if belowPoint.y < len(heightMap) && pointInBason(point, belowPoint, heightMap) {
				queue.enqueue(belowPoint)
			}
			if rightPoint.x < len(heightMap[0]) && pointInBason(point, rightPoint, heightMap) {
				queue.enqueue(rightPoint)
			}
		}
	}
	return bason
}

func pointInBason(mainPoint Point, checkPoint Point, heightMap [][]int) bool {
	if heightMap[checkPoint.y][checkPoint.x] == 9 {
		return false
	}
	if heightMap[checkPoint.y][checkPoint.x] == heightMap[mainPoint.y][mainPoint.x] || heightMap[checkPoint.y][checkPoint.x] > heightMap[mainPoint.y][mainPoint.x] {
		return true
	}
	return false
}

type Node struct {
	val  *Point
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
	len  int
}

func (q *Queue) enqueue(p Point) {
	var n Node
	n.val = &p
	if q.len == 0 {
		q.head = &n
		q.tail = &n
	} else {
		tail := q.tail
		tail.next = &n
		q.tail = &n
	}
	q.len++
}

func (q *Queue) dequeue() Point {
	head := q.head
	if head == q.tail {
		q.head = nil
		q.tail = nil
	} else {
		q.head = head.next
	}
	q.len--
	return *head.val
}
