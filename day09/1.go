package main

func IsNumberSumOfTwo(pool []int, value int) bool {
	for i := range pool {
		for j := i + 1; j < len(pool); j++ {
			if pool[i]+pool[j] == value {
				return true
			}
		}
	}
	return false
}

func FindFirstNumberNotSumOf(values []int, preambleSize int) int {
	for i := range values {
		value := values[preambleSize+i]
		if IsNumberSumOfTwo(values[i:preambleSize+i], value) == false {
			return value
		}
	}
	return -1
}
