package part3_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nasjp-sandbox/algorithms-and-data-structures/part3"
)

func TestShellSort(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   []int
		out  *part3.ShellSortOut
	}{
		{
			"case1",
			[]int{5, 1, 4, 3, 2},
			&part3.ShellSortOut{
				Out: []int{1, 2, 3, 4, 5},
				G:   []int{4, 1},
				Cnt: 3,
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if want, got := tt.out, part3.ShellSort(tt.in); !cmp.Equal(want, got) {
				t.Errorf("mismatch (-want +got):\n%s", cmp.Diff(want, got))
			}
		})
	}
}
