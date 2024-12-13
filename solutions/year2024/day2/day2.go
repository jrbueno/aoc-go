package day2

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

func RunAll(year int, day int, input []string, input2 []string) (partOneResult any, partTwoResult any) {
	return RunPartOne(input, year, day), RunPartTwo(input2, year, day)
}

func RunPartOne(input []string, year int, day int) any {
	safeCount := 0
	for _, l := range input {
		var r []int
		nums := strings.Fields(l)
		isSafe := true
		direction := 0
		for i, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("Error parsing number %s\n%v", num, err)
			}
			r = append(r, n)

			if i == 0 {
				continue
			}
			if direction > 0 && r[i-1] < r[i] {
				isSafe = false
				break
			}
			if direction < 0 && r[i-1] > r[i] {
				isSafe = false
				break
			}
			direction = r[i-1] - r[i]
			d := int(math.Abs(float64(r[i-1] - r[i])))
			if d == 0 || d > 3 {
				isSafe = false
				break
			}
		}
		if isSafe {
			safeCount++
		}
		//fmt.Printf("Row: %v Safe:%v\n", r, isSafe)
	}

	fmt.Printf("Total Safe: %d\n", safeCount)
	return safeCount
}

func RunPartTwo(input []string, year int, day int) any {
	safeCount := 0
	for _, l := range input {
		var r []int
		reports := strings.Fields(l)
		var isSafe bool
		for _, num := range reports {
			n, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("Error parsing number %s\n%v", num, err)
			}
			r = append(r, n)
		}

		//fmt.Printf("Processing Row: %v\n", r)
		isSafe = compareLevel(r)
		if !isSafe {
			fmt.Printf("UnSafe Row: %v\n", r)

			for i := 0; i < len(r); i++ {
				blah := slices.Delete(slices.Clone(r), i, i+1)
				fmt.Printf("Blah Row: %v\n", blah)
				isSafe = compareLevel(blah)
				fmt.Printf("Blah Row Safe: %v\n", isSafe)
				if isSafe {
					break
				}
			}
		}
		if isSafe {
			fmt.Printf("Row: %v Safe:%v\n", r, isSafe)
			safeCount++
		}
	}

	fmt.Printf("Total Safe: %d\n", safeCount)
	return safeCount
}

func compareLevel(r []int) bool {

	isSafe := true
	direction := 0
	for i := 1; i < len(r); i++ {
		if direction > 0 && r[i-1] < r[i] {
			isSafe = false
			break
		}
		if direction < 0 && r[i-1] > r[i] {
			isSafe = false
			break
		}
		d := int(math.Abs(float64(r[i-1] - r[i])))
		if d == 0 || d > 3 {
			isSafe = false
			break
		}
		direction = r[i-1] - r[i]
	}
	return isSafe
}
