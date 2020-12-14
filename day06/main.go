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
	groups := strings.Split(string(content), "\n\n")

	counter := 0
	counter2 := 0

	for _, group := range groups {
		counter += GetNumberOfQuestionsAnsweredYes(group)
		counter2 += GetNumberOfQuestionsAnsweredYesByEveryone(group)
	}

	println("Part 1")
	println(counter)

	println("Part 2")
	println(counter2)
}
