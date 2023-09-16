package main

import (
	"fmt"
	"github.com/samallen659/advent_of_code/2021/pkg/days/day1"
	"github.com/samallen659/advent_of_code/2021/pkg/days/day10"
	"github.com/samallen659/advent_of_code/2021/pkg/days/day11"
	"github.com/samallen659/advent_of_code/2021/pkg/days/day12"
	"github.com/samallen659/advent_of_code/2021/pkg/days/day13"
	"github.com/samallen659/advent_of_code/2021/pkg/days/day14"
	"github.com/samallen659/advent_of_code/2021/pkg/days/day15"
	"github.com/samallen659/advent_of_code/2021/pkg/days/day2"
	"github.com/samallen659/advent_of_code/2021/pkg/days/day3"
	"github.com/samallen659/advent_of_code/2021/pkg/days/day4"
	"github.com/samallen659/advent_of_code/2021/pkg/days/day5"
	"github.com/samallen659/advent_of_code/2021/pkg/days/day6"
	"github.com/samallen659/advent_of_code/2021/pkg/days/day7"
	"github.com/samallen659/advent_of_code/2021/pkg/days/day8"
	"github.com/samallen659/advent_of_code/2021/pkg/days/day9"
	"os"
)

func main() {
	dayArgument := os.Args[1:][0]
	switch dayArgument {
	case "1":
		fmt.Printf("%d\n", day1.Part1())
		fmt.Printf("%d\n", day1.Part2())
	case "2":
		fmt.Printf("%d\n", day2.Part1())
		fmt.Printf("%d\n", day2.Part2())
	case "3":
		fmt.Printf("%d\n", day3.Part1())
		fmt.Printf("%d\n", day3.Part2())
	case "4":
		fmt.Printf("%d\n", day4.Part1())
		fmt.Printf("%d\n", day4.Part2())
	case "5":
		fmt.Printf("%d\n", day5.Part1())
		fmt.Printf("%d\n", day5.Part2())
	case "6":
		fmt.Printf("%d\n", day6.Part1())
		fmt.Printf("%d\n", day6.Part2())
	case "7":
		fmt.Printf("%d\n", day7.Part1())
		fmt.Printf("%d\n", day7.Part2())
	case "8":
		fmt.Printf("%d\n", day8.Part1())
		fmt.Printf("%d\n", day8.Part2())
	case "9":
		fmt.Printf("%d\n", day9.Part1())
		fmt.Printf("%d\n", day9.Part2())
	case "10":
		fmt.Printf("%d\n", day10.Part1())
		fmt.Printf("%d\n", day10.Part2())
	case "11":
		fmt.Printf("%d\n", day11.Part1())
		fmt.Printf("%d\n", day11.Part2())
	case "12":
		fmt.Printf("%d\n", day12.Part1())
		fmt.Printf("%d\n", day12.Part2())
	case "13":
		fmt.Printf("%d\n", day13.Part1())
		day13.Part2()
	case "14":
		fmt.Printf("%d\n", day14.Part1())
		fmt.Printf("%d\n", day14.Part2())
	case "15":
		fmt.Printf("%d\n", day15.Part1())
	}
}
