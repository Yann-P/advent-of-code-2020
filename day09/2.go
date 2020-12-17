package main

func FindContiguousValuesThatSumTo(values []int, goalSum int) []int {
	position := 0
	sum := 0
	vals := []int{}

	for i := 0; i < len(values); i++ {

		sum += values[i]
		vals = append(vals, values[i])

		if sum == goalSum {
			return vals
		}

		if sum > goalSum {
			position++
			i = position
			sum = 0
			vals = []int{}
		}

	}

	return []int{}
}
