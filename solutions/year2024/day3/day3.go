package day3

import (
	"fmt"
	"regexp"
	"strconv"
)

func RunPartOne(lines []string, year int, day int) any {
	results := 0

	for _, line := range lines {
		re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			fmt.Printf("Match: %v\n", match)
			a, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Printf("Error converting %v to int: %v\n", match[1], err)
				continue
			}
			b, err := strconv.Atoi(match[2])
			if err != nil {
				fmt.Printf("Error converting %v to int: %v\n", match[2], err)
				continue
			}
			results += a * b
			//fmt.Printf("a: %v, b: %v\n", a, b)
		}

	}
	return results
}

func RunAll(year int, day int, input []string, input2 []string) (partOneResult any, partTwoResult any) {
	return RunPartOne(input, year, day), RunPartTwo(input2, year, day)
}

func RunPartTwo(input []string, year int, day int) any {
	return 0
}
