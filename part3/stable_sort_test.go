package part3_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nasjp-sandbox/algorithms-and-data-structures/part3"
)

func TestStableSort(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   []string
		out  []*part3.StableSortOut
	}{
		{
			"case1",
			[]string{"H4", "C9", "S4", "D2", "C3"},
			[]*part3.StableSortOut{
				{Out: []string{"D2", "C3", "H4", "S4", "C9"}, Stable: true},
				{Out: []string{"D2", "C3", "S4", "H4", "C9"}, Stable: false},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if want, got := tt.out, part3.StableSort(tt.in); !cmp.Equal(want, got) {
				t.Errorf("mismatch (-want +got):\n%s", cmp.Diff(want, got))
			}
		})
	}
}
