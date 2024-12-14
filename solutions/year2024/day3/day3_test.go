package day3

import (
	"testing"
)

func TestDay3(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		lines := []string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"}
		want := 161
		got := RunPartOne(lines, 2024, 3)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
