package part3_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nasjp-sandbox/algorithms-and-data-structures/part3"
)

func TestInsertionSort(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   []int
		out  [][]int
	}{
		{
			"case1",
			[]int{5, 2, 4, 6, 1, 3},
			[][]int{
				{5, 2, 4, 6, 1, 3},
				{2, 5, 4, 6, 1, 3},
				{2, 4, 5, 6, 1, 3},
				{2, 4, 5, 6, 1, 3},
				{1, 2, 4, 5, 6, 3},
				{1, 2, 3, 4, 5, 6},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if want, got := tt.out, part3.InsertionSort(tt.in); !cmp.Equal(want, got) {
				t.Errorf("mismatch (-want +got):\n%s", cmp.Diff(want, got))
			}
		})
	}
}
