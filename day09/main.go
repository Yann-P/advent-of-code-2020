package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	numbers := []int{}

	for _, val := range lines {
		if val == "" {
			continue
		}
		num, _ := strconv.Atoi(val)
		numbers = append(numbers, num)
	}

	println("Part 1")

	key := FindFirstNumberNotSumOf(numbers, 25)
	println(key)

	println("Part 2")

	ans := FindContiguousValuesThatSumTo(numbers, key)

	min, max := MinMax(ans)

	println(min + max)

}
