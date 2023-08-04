package day4

import (
	// "fmt"
	"strconv"
	"strings"

	"github.com/samallen659/advent_of_code/2021/pkg/utils"
)

var input = utils.ReadInput("/inputs/day4/input1.txt")
var data = strings.Split(input, "\n\n")
var moves = strings.Split(data[0], ",")
var boardData = data[1:]

func Part1() int {
	var boards []board
	for _, bd := range boardData {
		boards = append(boards, createBoard(bd))
	}
	var winner board
	var finalMove int
turn:
	for _, move := range moves {
		for _, b := range boards {
			moveNum, _ := strconv.Atoi(move)
			b.MarkNumber(moveNum)
			if b.CheckWin() {
				winner = b
				finalMove = moveNum
				break turn
			}
		}
	}
	return finalMove * sumUnmakedNumberes(winner)
}

func Part2() int {
	var boards []board
	for _, bd := range boardData {
		boards = append(boards, createBoard(bd))
	}
	boardMap := make(map[int]board)
	for i, b := range boards {
		boardMap[i] = b
	}
	var loser board
	var finalMove int
turn:
	for _, move := range moves {
		moveNum, _ := strconv.Atoi(move)
		for k, v := range boardMap {
			v.MarkNumber(moveNum)
			if v.CheckWin() {
				loser = v
				finalMove = moveNum
				delete(boardMap, k)
				if len(boardMap) == 0 {
					break turn
				}
			}
		}
	}
	return finalMove * sumUnmakedNumberes(loser)

}

func sumUnmakedNumberes(b board) int {
	sum := 0
	for key := range b.left {
		sum += key
	}
	return sum
}

func createBoard(data string) board {
	boardData := strings.Split(data, "\n")
	colLen := len(boardData)
	rowLen := len(boardData[0])
	var newBoard board
	newBoard.left = make(map[int]bool)
	for i := 0; i < colLen; i++ {
		rowIndex := 0
		for j := 2; j < rowLen+1; j++ {
			if (j+1)%3 == 0 {
				number, _ := strconv.Atoi(strings.Replace(boardData[i][j-2:j], " ", "0", 1))
				newBoard.tiles[i][rowIndex] = number
				newBoard.left[number] = true
				rowIndex++
			}
		}
	}
	return newBoard
}

type board struct {
	tiles [5][5]int
	left  map[int]bool
	won   bool
}

func (b board) MarkNumber(num int) {
	if b.left[num] {
		delete(b.left, num)
	}
}

func (b board) CheckWin() bool {
	for i := 0; i < 5; i++ {
		row := 0
		for j := 0; j < 5; j++ {
			if b.left[b.tiles[i][j]] {
				break
			} else {
				row++
			}
		}
		if row == 5 {
			return true
		}
		col := 0
		for j := 0; j < 5; j++ {
			if b.left[b.tiles[j][i]] {
				break
			} else {
				col++
			}
		}
		if col == 5 {
			return true
		}
	}
	return false
}
