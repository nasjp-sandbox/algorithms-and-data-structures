package part3_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nasjp-sandbox/algorithms-and-data-structures/part3"
)

func TestBubleSort(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   []int
		out  *part3.BubleSortOut
	}{
		{
			"case1",
			[]int{5, 3, 2, 4, 1},
			&part3.BubleSortOut{
				Out:       []int{1, 2, 3, 4, 5},
				SwapCount: 8,
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if want, got := tt.out, part3.BubleSort(tt.in); !cmp.Equal(want, got) {
				t.Errorf("mismatch (-want +got):\n%s", cmp.Diff(want, got))
			}
		})
	}
}
