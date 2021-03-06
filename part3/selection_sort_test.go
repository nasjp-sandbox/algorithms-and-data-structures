package part3_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nasjp-sandbox/algorithms-and-data-structures/part3"
)

func TestSelectionSort(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   []int
		out  *part3.SelectionSortOut
	}{
		{
			"case1",
			[]int{5, 6, 4, 2, 1, 3},
			&part3.SelectionSortOut{
				Out:       []int{1, 2, 3, 4, 5, 6},
				SwapCount: 4,
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if want, got := tt.out, part3.SelectionSort(tt.in); !cmp.Equal(want, got) {
				t.Errorf("mismatch (-want +got):\n%s", cmp.Diff(want, got))
			}
		})
	}
}
