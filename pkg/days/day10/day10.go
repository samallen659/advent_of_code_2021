package day10

import (
	"github.com/samallen659/advent_of_code/2021/pkg/utils"
	"math"
	"sort"
	"strings"
)

type Node struct {
	val  *string
	next *Node
}

type Stack struct {
	head *Node
	len  int
}

func (s *Stack) push(val string) {
	var n Node
	n.val = &val
	if s.len == 0 {
		s.head = &n
	} else {
		head := s.head
		n.next = head
		s.head = &n
	}
	s.len++
}

func (s *Stack) pop() *string {
	if s.len == 0 {
		return nil
	}
	head := s.head
	if s.len > 1 {
		s.head = head.next
	}
	s.len--
	return head.val
}

func (s *Stack) print() string {
	var loop func(head *Node) string
	loop = func(head *Node) string {
		if head.next != nil {
			return loop(head.next) + *head.val
		}
		return *head.val
	}

	return loop(s.head)
}

var input = utils.ReadInput("/inputs/day10/input1.txt")
var data = strings.Split(input, "\n")

func Part1() int {
	openingChunks := "({[<"
	var corruptedChunks []string
	for _, line := range data {
		var s Stack
		chars := strings.Split(line, "")
	charLoop:
		for _, char := range chars {
			if strings.Contains(openingChunks, char) {
				s.push(char)
			} else {
				currentChunk := s.pop()
				switch char {
				case ")":
					if *currentChunk != "(" {
						corruptedChunks = append(corruptedChunks, char)
						break charLoop
					}
				case "}":
					if *currentChunk != "{" {
						corruptedChunks = append(corruptedChunks, char)
						break charLoop
					}
				case "]":
					if *currentChunk != "[" {
						corruptedChunks = append(corruptedChunks, char)
						break charLoop
					}
				case ">":
					if *currentChunk != "<" {
						corruptedChunks = append(corruptedChunks, char)
						break charLoop
					}
				}
			}
		}
	}
	score := 0
	for _, char := range corruptedChunks {
		switch char {
		case ")":
			score += 3
		case "]":
			score += 57
		case "}":
			score += 1197
		case ">":
			score += 25137
		}
	}
	return score
}

func Part2() int {
	openingChunks := "({[<"
	var lineAutocompleteScores []int
	for _, line := range data {
		var s Stack
		chars := strings.Split(line, "")
		isCorrupted := false
	charLoop:
		for _, char := range chars {
			if strings.Contains(openingChunks, char) {
				s.push(char)
			} else {
				currentChunk := s.pop()
				switch char {
				case ")":
					if *currentChunk != "(" {
						isCorrupted = true
						break charLoop
					}
				case "}":
					if *currentChunk != "{" {
						isCorrupted = true
						break charLoop
					}
				case "]":
					if *currentChunk != "[" {
						isCorrupted = true
						break charLoop
					}
				case ">":
					if *currentChunk != "<" {
						isCorrupted = true
						break charLoop
					}
				}
			}
		}
		if !isCorrupted {
			score := 0
			for s.len > 0 {
				score *= 5
				currentChunk := s.pop()
				switch *currentChunk {
				case "(":
					score += 1
				case "[":
					score += 2
				case "{":
					score += 3
				case "<":
					score += 4
				}
			}
			lineAutocompleteScores = append(lineAutocompleteScores, score)
		}
	}
	sort.Ints(lineAutocompleteScores)
	return lineAutocompleteScores[int(math.Round(float64(len(lineAutocompleteScores))/2))-1]
}
