package main

import (
	"regexp"
	"strconv"
	"strings"
)

func IsPasswordValid(s string) bool {
	r, _ := regexp.Compile("^(\\d+)-(\\d+) ([a-z]{1}): ([a-z]+)$")

	extracted := r.FindStringSubmatch(s)

	if len(extracted) != 5 {
		return false
	}

	min, _ := strconv.Atoi(extracted[1])
	max, _ := strconv.Atoi(extracted[2])
	letter := extracted[3]
	password := extracted[4]
	count := strings.Count(password, letter)

	if count < min || count > max {
		return false
	}

	return true
}
