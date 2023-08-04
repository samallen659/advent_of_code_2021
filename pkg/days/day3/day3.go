package day3

import (
	"strconv"
	"strings"

	"github.com/samallen659/advent_of_code/2021/pkg/utils"
)

func Part1() int64 {
	input := utils.ReadInput("/inputs/day3/input1.txt")
	data := strings.Split(input, "\n")
	gammaStr := getGamma(data)
	epsilonStr := getEpsilon(gammaStr)
	gamma, _ := strconv.ParseInt(gammaStr, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonStr, 2, 64)
	return gamma * epsilon
}

func Part2() int64 {
	input := utils.ReadInput("/inputs/day3/input1.txt")
	data := strings.Split(input, "\n")
	oxygenGenerator, _ := strconv.ParseInt(getOxygenGenerator(data), 2, 64)
	c02Generator, _ := strconv.ParseInt(getCO2Generator(data), 2, 64)
	return oxygenGenerator * c02Generator
}

func getGamma(data []string) string {
	lineLen := len(data[0])
	gamma := ""
	for i := 0; i < lineLen; i++ {
		oneCount := 0
		zeroCount := 0
		for j := 0; j < len(data); j++ {
			if strings.Split(data[j], "")[i] == "1" {
				oneCount++
				continue
			}
			zeroCount++
		}
		if oneCount > zeroCount {
			gamma = gamma + "1"
			continue
		}
		gamma = gamma + "0"
	}
	return gamma
}

func getEpsilon(gamma string) string {
	epsilon := ""
	for _, s := range gamma {
		if string(s) == "1" {
			epsilon = epsilon + "0"

		}
		epsilon = epsilon + "1"
	}
	return epsilon
}

func getOxygenGenerator(data []string) string {
	bitLen := len(data[0])
	pos := bitLen - 1
	o2_rem := data
	for pos >= 0 && len(o2_rem) > 1 {
		var onesFiltered []string
		var zerosFiltered []string
		for i := 0; i < len(o2_rem); i++ {
			binNumber, _ := strconv.ParseInt(o2_rem[i], 2, 64)
			if binNumber&(1<<pos) == (1 << pos) {
				onesFiltered = append(onesFiltered, o2_rem[i])
			} else {
				zerosFiltered = append(zerosFiltered, o2_rem[i])
			}
		}
		if len(onesFiltered) >= len(zerosFiltered) {
			o2_rem = onesFiltered
		} else {
			o2_rem = zerosFiltered
		}
		pos--
	}
	return o2_rem[0]
}
func getCO2Generator(data []string) string {
	bitLen := len(data[0])
	pos := bitLen - 1
	co2_rem := data
	for pos >= 0 && len(co2_rem) > 1 {
		var onesFiltered []string
		var zerosFiltered []string
		for i := 0; i < len(co2_rem); i++ {
			binNumber, _ := strconv.ParseInt(co2_rem[i], 2, 64)
			if binNumber&(1<<pos) == (1 << pos) {
				onesFiltered = append(onesFiltered, co2_rem[i])
			} else {
				zerosFiltered = append(zerosFiltered, co2_rem[i])
			}
		}
		if len(onesFiltered) < len(zerosFiltered) {
			co2_rem = onesFiltered
		} else {
			co2_rem = zerosFiltered
		}
		pos--
	}
	return co2_rem[0]
}
