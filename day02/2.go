package main

import (
	"regexp"
	"strconv"
)

func IsPasswordValid_2(s string) bool {
	r, _ := regexp.Compile("^(\\d+)-(\\d+) ([a-z]{1}): ([a-z]+)$")

	extracted := r.FindStringSubmatch(s)

	if len(extracted) != 5 {
		return false
	}

	pos1, _ := strconv.Atoi(extracted[1]) // starts at 1
	pos2, _ := strconv.Atoi(extracted[2])
	letter := extracted[3][0]
	password := extracted[4]

	if pos1 < 1 || pos2 < 1 || pos1 > len(password) || pos2 > len(password) {
		return false
	}

	pos1Match := password[pos1-1] == letter
	pos2Match := password[pos2-1] == letter

	if pos1Match == pos2Match { // both true or both false, we want a xor
		return false
	}

	return true
}
