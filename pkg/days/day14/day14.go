package day14

import (
	"fmt"
	"github.com/samallen659/advent_of_code/2021/pkg/utils"
	"strings"
)

var input = utils.ReadInput("/inputs/day14/input1.txt")
var data = strings.Split(input, "\n\n")

func Part1() int {
	template := data[0]
	pairMap := buildPairInsertionMap(data[1])

	polymer := runPolymerizationStep(10, template, pairMap)
	occurenceMap := buildOccurenceMap(polymer)

	var maxOccurence string
	var minOccurence string
	for k, v := range occurenceMap {
		if minOccurence == "" {
			minOccurence = k
		}
		if v > occurenceMap[maxOccurence] {
			maxOccurence = k
		}
		if v < occurenceMap[minOccurence] {
			minOccurence = k
		}
	}

	result := occurenceMap[maxOccurence] - occurenceMap[minOccurence]
	return int(result)
}

func Part2() int64 {
	template := data[0]
	pairInsertMap := buildPairInsertionMap(data[1])
	pairMap := make(map[string]int64)
	for i := 1; i < len(template); i++ {
		str := fmt.Sprintf("%s%s", string(template[i-1]), string(template[i]))
		pairMap[str] += 1
	}

	runPolymerizationStep2(40, &pairInsertMap, &pairMap)
	occurenceMap := buildOccurenceMap2(template, &pairMap)

	var maxOccurence string
	var minOccurence string
	for k, v := range occurenceMap {
		if minOccurence == "" {
			minOccurence = k
		}
		if v > occurenceMap[maxOccurence] {
			maxOccurence = k
		}
		if v < occurenceMap[minOccurence] {
			minOccurence = k
		}
	}

	result := occurenceMap[maxOccurence] - occurenceMap[minOccurence]
	return result
}

func buildPairInsertionMap(pairInsertion string) map[string]string {
	pairInsertionLines := strings.Split(pairInsertion, "\n")
	pairInsertionMap := make(map[string]string)

	for _, pair := range pairInsertionLines {
		pairRule := strings.Split(pair, " -> ")
		pairInsertionMap[pairRule[0]] = pairRule[1]
	}

	return pairInsertionMap
}

func runPolymerizationStep2(steps int, pairInsertMap *map[string]string, pairMap *map[string]int64) {
	for i := 0; i < steps; i++ {
		tmpMap := make(map[string]int64)
		for k, v := range *pairMap {
			insert := (*pairInsertMap)[k]
			p1 := string(k[0]) + insert
			p2 := insert + string(k[1])
			tmpMap[p1] += v
			tmpMap[p2] += v
		}
		*pairMap = tmpMap
	}
}

func runPolymerizationStep(steps int, template string, pairMap map[string]string) string {
	temp := template
	for i := 0; i < steps; i++ {
		newTemp := ""
		for j := 1; j < len(temp); j++ {
			insert := pairMap[temp[j-1:j+1]]
			newTemp += fmt.Sprintf("%s%s", string(temp[j-1]), insert)
		}
		newTemp += string(temp[len(temp)-1])

		temp = newTemp
	}

	return temp
}

func buildOccurenceMap2(template string, pairMap *map[string]int64) map[string]int64 {
	occurenceMap := make(map[string]int64)
	occurenceMap[string(template[0])] += 1
	occurenceMap[string(template[len(template)-1])] += 1

	for k, v := range *pairMap {
		occurenceMap[string(k[0])] += v
		occurenceMap[string(k[1])] += v
	}

	for k, v := range occurenceMap {
		occurenceMap[k] = v / 2
	}

	return occurenceMap
}

func buildOccurenceMap(polymer string) map[string]int64 {
	occurenceMap := make(map[string]int64)
	for _, s := range polymer {
		occurenceMap[string(s)] += 1
	}

	return occurenceMap
}
