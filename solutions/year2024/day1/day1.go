package day1

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

func RunPartOne(input []string, year int, day int) any {
	//fmt.Printf("Running Part One solution for year %d day %d\n", year, day)
	left, right := parseInput(input)
	if left == nil || right == nil {
		log.Printf("Error parsing the file - left or right arrays are nil")
		return 0
	}
	//sort the arrays
	slices.Sort(left)
	slices.Sort(right)

	//verify the arrays are the same length
	if len(left) != len(right) {
		log.Fatalf("Error parsing the file - left and right arrays are not the same length")
	}

	//print the arrays
	//fmt.Printf("Left Side: %v\n", left)
	//fmt.Printf("Right Side: %v\n", right)

	//calculate the total distance
	var totalDistance int
	for i := 0; i < len(left); i++ {
		distance := int(math.Abs(float64(left[i] - right[i])))
		//fmt.Println(distance)
		totalDistance += distance
	}
	//fmt.Printf("Total Distance: %v\n", totalDistance)
	return totalDistance
}

func RunPartTwo(input []string, year int, day int) any {
	//fmt.Printf("Running Part Two solution for year %d day %d\n", year, day)
	left, right := parseInput(input)
	if left == nil || right == nil {
		log.Printf("Error parsing the file - left or right arrays are nil")
		return 0
	}
	//sort the arrays
	slices.Sort(left)
	slices.Sort(right)

	//verify the arrays are the same length
	if len(left) != len(right) {
		log.Fatalf("Error parsing the file - left and right arrays are not the same length")
	}

	//iterate through the right array and count the number of times each number appears
	numCounts := make(map[int]int)
	for _, num := range right {
		numCounts[num]++
	}
	//calculate the similarity score
	var similarityScore int
	for _, num := range left {
		if numCounts[num] > 0 {
			similarityScore += num * numCounts[num]
		}
	}
	return similarityScore
}

func RunAll(year int, day int, partOneInput []string, partTwoInput []string) (any, any) {
	fmt.Printf("Running All solution for year %d day %d\n", year, day)
	partOneResult := RunPartOne(partOneInput, year, day)
	partTwoResult := RunPartTwo(partTwoInput, year, day)
	return partOneResult, partTwoResult
}

func parseInput(lines []string) ([]int, []int) {
	//check if the file is nil
	if len(lines) == 0 {
		log.Printf("No Data to Parse")
		return nil, nil
	}
	var left, right []int
	scanner := bufio.NewScanner(strings.NewReader(strings.Join(lines, "\n")))
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		if len(words) != 2 {
			log.Fatalf("Error parsing the file - invalid line: %v", line)
		}
		if n, err := strconv.Atoi(words[0]); err == nil {
			left = append(left, n)
		} else {
			log.Fatalf("Error parsing the file: %v", err)
		}
		if n, err := strconv.Atoi(words[1]); err == nil {
			right = append(right, n)
		} else {
			log.Fatalf("Error parsing the file: %v", err)
		}
	}
	if err := scanner.Err(); err != nil {
		// handle error
		log.Fatalf("Error parsing the file: %v", err)
	}
	return left, right
}
