package main

func MultiplyNumbersThatSumTo2020_2(input []int) int {
	n := len(input)
	for i := range input {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if input[i]+input[j]+input[k] == 2020 {
					return input[i] * input[j] * input[k]
				}
			}
		}
	}
	return -1
}
