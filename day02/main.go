package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")

	validCount := 0
	validCount2 := 0

	for _, password := range lines {
		if IsPasswordValid(password) {
			validCount++
		}
		if IsPasswordValid_2(password) {
			validCount2++
		}
	}

	fmt.Println(validCount)
	fmt.Println(validCount2)

}
