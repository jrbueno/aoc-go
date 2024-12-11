package day2

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func RunAll(year int, day int, input []string, input2 []string) (partOneResult any, partTwoResult any) {
	return RunPartOne(input, year, day), RunPartTwo(input2, year, day)
}

func RunPartOne(input []string, year int, day int) any {
	//var reports []bool
	safeCount := 0
	scanner := bufio.NewScanner(strings.NewReader(strings.Join(input, "\n")))
	for scanner.Scan() {
		var r []int
		l := scanner.Text()
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
		fmt.Printf("Row: %v Safe:%v\n", r, isSafe)
		//reports = append(reports, isSafe)
	}

	fmt.Printf("Total Safe: %d\n", safeCount)
	return safeCount
}

func RunPartTwo(input []string, year int, day int) any {
	return nil
}
