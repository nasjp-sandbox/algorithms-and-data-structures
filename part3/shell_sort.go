package part3

type ShellSortOut struct {
	Out []int
	G   []int
	Cnt int
}

func ShellSort(in []int) *ShellSortOut {
	g := make([]int, 0)

	for h := 1; ; h = 3*h + 1 {
		if h > len(in) {
			break
		}

		g = append([]int{h}, g...)
	}

	out := make([]int, len(in))
	copy(out, in)

	var cnt int

	for _, gg := range g {
		var c int
		out, c = insertionSort(out, gg)
		cnt += c
	}

	return &ShellSortOut{
		Out: out,
		Cnt: cnt,
		G:   g,
	}
}

func insertionSort(in []int, g int) ([]int, int) {
	out := make([]int, len(in))
	copy(out, in)

	var cnt int

	for i, v := range out {
		j := i - g

		for j >= 0 && out[j] > v {
			out[j+g] = out[j]
			j -= g
			cnt++
		}

		out[j+g] = v
	}

	return out, cnt
}
