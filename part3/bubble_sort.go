package part3

type BubbleSortOut struct {
	Out       []int
	SwapCount int
}

func BubbleSort(in []int) *BubbleSortOut {
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

	return &BubbleSortOut{
		Out:       in,
		SwapCount: swap,
	}
}
