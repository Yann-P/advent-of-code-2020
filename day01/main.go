package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	numbers := []int{}

	for _, val := range lines {
		num, _ := strconv.Atoi(val)
		numbers = append(numbers, num)
	}
	println(MultiplyNumbersThatSumTo2020(numbers[:]))
	println(MultiplyNumbersThatSumTo2020_2(numbers[:]))
}
