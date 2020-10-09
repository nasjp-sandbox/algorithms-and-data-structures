package part3

import (
	"strconv"
)

type StableSortOut struct {
	Out    []string
	Stable bool
}

func StableSort(in []string) []*StableSortOut {
	bubbleResult := bubbleSort(in)

	selectionResult := selectionSort(in)

	out := []*StableSortOut{
		{Out: bubbleResult, Stable: isStable(in, bubbleResult)},
		{Out: selectionResult, Stable: isStable(in, selectionResult)},
	}

	return out
}

func bubbleSort(in []string) []string {
	out := make([]string, len(in))
	copy(out, in)

	for flag, i := true, 0; flag; {
		flag = false

		for j := len(out) - 1; j > i; j-- {
			jNum, _ := strconv.Atoi(string(out[j][1]))
			beforeJNum, _ := strconv.Atoi(string(out[j-1][1]))

			if jNum < beforeJNum {
				out[j], out[j-1] = out[j-1], out[j]
				flag = true
			}
		}
		i++
	}

	return out
}

func selectionSort(in []string) []string {
	out := make([]string, len(in))
	copy(out, in)

	for i := 0; i < len(out); i++ {
		minj := i

		for j := minj; j < len(out); j++ {
			minJNum, _ := strconv.Atoi(string(out[minj][1]))
			jNum, _ := strconv.Atoi(string(out[j][1]))

			if minJNum > jNum {
				minj = j
			}
		}

		if i != minj {
			out[i], out[minj] = out[minj], out[i]
		}
	}

	return out
}

func isStable(in []string, out []string) bool {
	for ini := 0; ini < len(in); ini++ {
		for inj := ini + 1; inj < len(in); inj++ {
			for outi := 0; outi < len(in); outi++ {
				for outj := outi + 1; outj < len(in); outj++ {
					if in[ini][1] == in[inj][1] && in[ini] == out[outj] && in[inj] == out[outi] {
						return false
					}
				}
			}
		}
	}

	return true
}
