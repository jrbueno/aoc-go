package day1

import (
	"testing"
)

func TestRunPartOne(t *testing.T) {
	t.Run("PartOne with nil file", func(t *testing.T) {
		got := RunPartOne(nil, 2024, 1)
		if got != nil {
			t.Errorf("got %v want %v", got, nil)
		}
	})

	t.Run("PartOne with empty slice", func(t *testing.T) {
		got := RunPartOne([]string{}, 2024, 1)
		if got != 0 {
			t.Errorf("got %v want %v", got, 0)
		}
	})
	t.Run("PartOne with empty slice", func(t *testing.T) {
		lines := []string{"3   4", "4   3", "2   5", "1   3", "3   9", "3   3"}
		got := RunPartOne(lines, 2024, 1)
		want := 11
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
