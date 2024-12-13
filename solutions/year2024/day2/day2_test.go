package day2

import (
	"testing"
)

func TestDay2PartOne(t *testing.T) {
	t.Run("PartOne with nil file", func(t *testing.T) {
		got := RunPartOne(nil, 2024, 2)
		if got != 0 {
			t.Errorf("got %v want %v", got, 0)
		}
	})
	t.Run("PartOne with test Input", func(t *testing.T) {
		lines := []string{"7 6 4 2 1", "1 2 7 8 9", "9 7 6 2 1", "1 3 2 4 5", "8 6 4 4 1", "1 3 6 7 9"}
		got := RunPartOne(lines, 2024, 2)
		want := 2
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("PartTwo with test Input", func(t *testing.T) {
		lines := []string{"7 6 4 2 1", "1 2 7 8 9", "9 7 6 2 1", "1 3 2 4 5", "8 6 4 4 1", "1 3 6 7 9"}
		got := RunPartTwo(lines, 2024, 2)
		want := 4
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("PartTwo with 59 61 59 61 63", func(t *testing.T) {
		lines := []string{"59 61 59 61 63"}
		got := RunPartTwo(lines, 2024, 2)
		want := 0
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
