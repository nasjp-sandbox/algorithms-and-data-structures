package part3

func InsertionSort(in []int) [][]int {
	outs := make([][]int, 0)

	first := make([]int, len(in))
	copy(first, in)

	outs = append(outs, first)

	for i := 1; i < len(in); i++ {
		out := make([]int, len(in))
		copy(out, outs[i-1])

		cur := out[i]

		j := i - 1

		for j >= 0 && out[j] > cur {
			out[j+1] = out[j]
			j--
		}

		out[j+1] = cur

		outs = append(outs, out)
	}

	return outs
}
