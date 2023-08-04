package day8

import (
	"github.com/samallen659/advent_of_code/2021/pkg/utils"
	"golang.org/x/exp/slices"
	"sort"
	"strconv"
	"strings"
)

var input = utils.ReadInput("/inputs/day8/input1.txt")
var data = strings.Split(input, "\n")
var uniqueSegmentNums = make(map[int]int)

func Part1() int {
	uniqueSegmentNums[7] = 8
	uniqueSegmentNums[3] = 7
	uniqueSegmentNums[4] = 4
	uniqueSegmentNums[2] = 1

	uniqueSegmentOutputCount := 0
	for _, d := range data {
		digitOutputStr := strings.Split(d, "|")
		digitOutput := strings.Split(digitOutputStr[1], " ")
		for _, digit := range digitOutput {
			_, ok := uniqueSegmentNums[len(digit)]
			if ok {
				uniqueSegmentOutputCount++
			}

		}
	}

	return uniqueSegmentOutputCount
}

func Part2() int {
	uniqueSegmentNums[7] = 8
	uniqueSegmentNums[3] = 7
	uniqueSegmentNums[4] = 4
	uniqueSegmentNums[2] = 1

	output := 0
	for _, d := range data {
		patternOutputStr := strings.Split(d, "|")
		signalOutput := strings.Split(patternOutputStr[0], " ")
		digitOutput := strings.Split(patternOutputStr[1], " ")
		var patterns []string

		for _, pattern := range signalOutput {
			if pattern != "" {
				p := strings.Split(pattern, "")
				sort.Strings(p)
				sortedP := strings.Join(p, "")
				patterns = append(patterns, sortedP)
			}
		}

		for _, pattern := range digitOutput {
			if pattern != "" {
				p := strings.Split(pattern, "")
				sort.Strings(p)
				sortedP := strings.Join(p, "")
				patterns = append(patterns, sortedP)
			}
		}

		var matchedPatterns [10]string
		matchedCount := 0
		for _, pattern := range patterns {
			v, unique := uniqueSegmentNums[len(pattern)]

			if unique && matchedPatterns[v] == "" {
				matchedPatterns[v] = pattern
				matchedCount++
			}
		}
		for matchedCount < 10 {
			for _, pattern := range patterns {
				if slices.Contains(matchedPatterns[:], pattern) {
					continue
				}
				if len(pattern) == 6 {
					if checkLenSix(pattern, &matchedPatterns) {
						matchedCount++
					}
				}
				if len(pattern) == 5 {
					if checkLenFive(pattern, &matchedPatterns) {
						matchedCount++
					}
				}
			}
		}
		numberStr := ""
		for _, pattern := range digitOutput {
			p := strings.Split(pattern, "")
			sort.Strings(p)
			sortedP := strings.Join(p, "")
			for j := 0; j < len(matchedPatterns); j++ {
				if sortedP == matchedPatterns[j] {
					numberStr += strconv.Itoa(j)
					break
				}
			}
		}
		number, _ := strconv.Atoi(numberStr)
		output += number
	}

	return output
}

func checkLenSix(pattern string, matchedPatterns *[10]string) bool {
	if (*matchedPatterns)[9] == "" && checkMatchedChars(pattern, (*matchedPatterns)[4]) {
		(*matchedPatterns)[9] = pattern
		return true
	}
	if (*matchedPatterns)[9] != "" && (*matchedPatterns)[0] == "" && checkMatchedChars(pattern, (*matchedPatterns)[7]) {
		(*matchedPatterns)[0] = pattern
		return true
	}
	if (*matchedPatterns)[9] != "" && (*matchedPatterns)[0] != "" {
		(*matchedPatterns)[6] = pattern
		return true
	}
	return false
}

func checkLenFive(pattern string, matchedPatterns *[10]string) bool {
	if (*matchedPatterns)[3] == "" && checkMatchedChars(pattern, (*matchedPatterns)[1]) {
		(*matchedPatterns)[3] = pattern
		return true
	}
	if (*matchedPatterns)[3] != "" && (*matchedPatterns)[9] != "" && (*matchedPatterns)[5] == "" && checkMatchedChars((*matchedPatterns)[9], pattern) {
		(*matchedPatterns)[5] = pattern
		return true
	}
	if (*matchedPatterns)[3] != "" && (*matchedPatterns)[5] != "" {
		(*matchedPatterns)[2] = pattern
		return true
	}
	return false
}

func checkMatchedChars(p1 string, p2 string) bool {
	a1 := strings.Split(p1, "")
	a2 := strings.Split(p2, "")
	if len(a1) > len(a2) {
		for _, a := range a2 {
			if !strings.Contains(p1, a) {
				return false
			}
		}
	} else {
		for _, a := range a1 {
			if !strings.Contains(p2, a) {
				return false
			}
		}
	}
	return true
}
