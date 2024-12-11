package day2

import (
	"testing"
)

func TestDay2PartOne(t *testing.T) {
	t.Run("PartOne with nil file", func(t *testing.T) {
		got := RunPartOne(nil, 2024, 2)
		if got != nil {
			t.Errorf("got %v want %v", got, nil)
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
}
