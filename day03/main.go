package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")

	println("Part 1")
	println(ComputeTreesEncountered(lines, 3, 1))

	println("Part 2")
	println(ComputeTreesEncountered(lines, 1, 1) * ComputeTreesEncountered(lines, 3, 1) * ComputeTreesEncountered(lines, 5, 1) * ComputeTreesEncountered(lines, 7, 1) * ComputeTreesEncountered(lines, 1, 2))
}
