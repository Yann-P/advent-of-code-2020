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
	passports := strings.Split(string(content), "\n\n")

	oneLinePassports := []string{}

	for _, passport := range passports {
		oneLinePassports = append(oneLinePassports, strings.ReplaceAll(passport, "\n", " "))
	}

	validPassports := 0
	strictlyValidPassports := 0

	for _, oneLinePassport := range oneLinePassports {
		if IsPassportValid(oneLinePassport) {
			validPassports++
		}
		if IsPassportValidStrict(oneLinePassport) {
			strictlyValidPassports++
		}
	}

	println("Part 1")
	println(validPassports)

	println("Part 2")
	println(strictlyValidPassports)

}
