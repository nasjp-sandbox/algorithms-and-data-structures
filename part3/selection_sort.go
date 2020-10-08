package part3

type SelectionSortOut struct {
	Out       []int
	SwapCount int
}

func SelectionSort(in []int) *SelectionSortOut {
	var swap int

	for i := 0; i < len(in); i++ {
		minj := i

		for j := minj; j < len(in); j++ {
			if in[minj] > in[j] {
				minj = j
			}
		}

		if i != minj {
			in[i], in[minj] = in[minj], in[i]
			swap++
		}
	}

	return &SelectionSortOut{
		Out:       in,
		SwapCount: swap,
	}
}
