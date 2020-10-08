package part3

type BubleSortOut struct {
	Out       []int
	SwapCount int
}

func BubleSort(in []int) *BubleSortOut {
	var swap int

	for flag, i := true, 0; flag; {
		flag = false

		for j := len(in) - 1; j > i; j-- {
			if in[j] < in[j-1] {
				in[j], in[j-1] = in[j-1], in[j]
				flag = true
				swap++
			}
		}
		i++
	}

	return &BubleSortOut{
		Out:       in,
		SwapCount: swap,
	}
}
