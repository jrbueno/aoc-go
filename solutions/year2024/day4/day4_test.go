package day4

import (
	"strings"
	"testing"
)

func TestDay4(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		got := RunPartOne(nil, 2024, 4)
		want := 0
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("Part 1 with Test Data", func(t *testing.T) {
		input := `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
		`
		lines := strings.Split(strings.TrimSpace(input), "\n")
		got := RunPartOne(lines, 2024, 4)
		want := 18
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
