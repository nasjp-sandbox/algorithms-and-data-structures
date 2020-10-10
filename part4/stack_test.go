package part4_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nasjp-sandbox/algorithms-and-data-structures/part4"
)

func TestStack(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   []string
		out  string
	}{
		{
			"case1",
			[]string{"1", "2", "+", "3", "4", "-", "*"},
			"-3",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if want, got := tt.out, part4.Stack(tt.in); !cmp.Equal(want, got) {
				t.Errorf("mismatch (-want +got):\n%s", cmp.Diff(want, got))
			}
		})
	}
}
