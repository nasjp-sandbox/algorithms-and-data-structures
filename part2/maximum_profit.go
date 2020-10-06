package part2

const minRate = 200000

func MaximumProfit(rates []int) int {
	max, min := -minRate, rates[0]

	for i := 1; i < len(rates); i++ {
		rate := rates[i]

		if gain := rate - min; max < gain {
			max = gain
		}

		if min > rate {
			min = rate
		}
	}

	return max
}
