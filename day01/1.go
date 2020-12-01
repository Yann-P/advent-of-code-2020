package main

func MultiplyNumbersThatSumTo2020(input []int) int {
	n := len(input)
	for i := range input {
		for j := i + 1; j < n; j++ {
			if input[i]+input[j] == 2020 {
				return input[i] * input[j]
			}
		}
	}
	return -1
}
