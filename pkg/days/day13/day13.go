package day13

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samallen659/advent_of_code/2021/pkg/utils"
)

var input = utils.ReadInput("/inputs/day13/input1.txt")
var data = strings.Split(input, "\n\n")
var pointStrings = strings.Split(data[0], "\n")
var foldsStrings = strings.Split(data[1], "\n")

type Graph [][]int

func (g *Graph) Fold(f Fold) {
	if f.direction == "x" {
		g.foldHorizontally(f.distance)
	} else {
		g.foldVertically(f.distance)
	}
}

func (g *Graph) foldVertically(distance int) {
	for i := distance + 1; i < len(*g); i++ {
		for j := 0; j < len((*g)[i]); j++ {
			if (*g)[i][j] == 1 {
				newPos := distance - (i - distance)
				(*g)[newPos][j] = 1
			}
		}
	}
	(*g) = (*g)[0:distance]
}

func (g *Graph) foldHorizontally(distance int) {
	for i := range *g {
		for j := distance + 1; j < len((*g)[i]); j++ {
			if (*g)[i][j] == 1 {
				newPos := distance - (j - distance)
				(*g)[i][newPos] = 1
			}
		}
		(*g)[i] = (*g)[i][0:distance]
	}
}

func (g *Graph) CountPoints() int {
	pointCount := 0
	for i := range *g {
		for j := range (*g)[0] {
			if (*g)[i][j] == 1 {
				pointCount++
			}
		}
	}

	return pointCount
}

type Fold struct {
	direction string
	distance  int
}

type Point struct {
	x int
	y int
}

func Part1() int {
	points, xLength, yLength := buildPointSlice(pointStrings)
	folds := buildFoldSlice(foldsStrings)
	graph := buildGraph(points, yLength, xLength)

	graph.Fold(folds[0])

	return graph.CountPoints()
}

func Part2() {
	points, xLength, yLength := buildPointSlice(pointStrings)
	folds := buildFoldSlice(foldsStrings)
	graph := buildGraph(points, yLength, xLength)

	for _, f := range folds {
		graph.Fold(f)
	}

	for _, g := range graph {
		fmt.Println(g)
	}
}

func buildPointSlice(pointStrings []string) ([]Point, int, int) {
	xLength := 0
	yLength := 0
	var points []Point
	for _, p := range pointStrings {
		pointSplit := strings.Split(p, ",")
		x, _ := strconv.Atoi(pointSplit[0])
		y, _ := strconv.Atoi(pointSplit[1])
		points = append(points, Point{x: x, y: y})
		if x > xLength {
			xLength = x
		}
		if y > yLength {
			yLength = y
		}
	}

	return points, xLength, yLength
}

func buildFoldSlice(foldStrings []string) []Fold {
	var folds []Fold
	for _, f := range foldsStrings {
		direction := f[11:12]
		distance, _ := strconv.Atoi(f[13:])
		folds = append(folds, Fold{direction: direction, distance: distance})
	}

	return folds
}

func buildGraph(points []Point, yLength int, xLength int) Graph {
	graph := make([][]int, yLength+1)
	for i := range graph {
		graph[i] = make([]int, xLength+1)
	}

	for _, p := range points {
		graph[p.y][p.x] = 1
	}

	return graph
}
