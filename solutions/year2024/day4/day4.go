package day4

func RunAll(year int, day int, input []string, input2 []string) (partOneResult any, partTwoResult any) {
	return RunPartOne(input, year, day), RunPartTwo(input2, year, day)
}

func RunPartOne(input []string, year int, day int) any {
	if input == nil {
		return 0
	}

	var directions = [][]int{
		{-1, 0},  // up
		{1, 0},   // down
		{0, -1},  // left
		{0, 1},   // right
		{-1, -1}, // up-left
		{-1, 1},  // up-right
		{1, -1},  // down-left
		{1, 1},   // down-right
	}
	foundCount := 0
	var word = []rune("XMAS")
	var wordmap = make([][]rune, len(input))
	for i, line := range input {
		wordmap[i] = []rune(line)
	}
	for r := 0; r < len(wordmap); r++ {
		for c := 0; c < len(wordmap[r]); c++ {
			for _, direction := range directions {
				rowIndex := direction[0]
				colIndex := direction[1]
				foundXmas := true
				for i := 0; i < len(word); i++ {
					rowOffset := r + (rowIndex * i)
					colOffset := c + (colIndex * i)
					if rowOffset < 0 || rowOffset >= len(wordmap) || colOffset < 0 || colOffset >= len(wordmap[r]) {
						foundXmas = false
						break
					}
					if wordmap[rowOffset][colOffset] != word[i] {
						foundXmas = false
						break
					}
				}
				if foundXmas {
					foundCount++
				}
			}
		}
	}
	return foundCount
}

func RunPartTwo(input []string, year int, day int) any {
	return nil
}
