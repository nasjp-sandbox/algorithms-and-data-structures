package part2_test

import (
	"testing"

	"github.com/nasjp-sandbox/algorithms-and-data-structures/part2"
)

func TestMaximumProfit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   []int
		out  int
	}{
		{"case1", []int{5, 3, 1, 3, 4, 3}, 3},
		{"case2", []int{4, 3, 2}, -1},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if want, got := tt.out, part2.MaximumProfit(tt.in); want != got {
				t.Errorf("want %d, but got %d", want, got)
			}
		})
	}

}
