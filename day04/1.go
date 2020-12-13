package main

import "strings"

func IsPassportValid(passport string) bool {

	fields := strings.Split(passport, " ")

	data := make(map[string]string)

	for _, field := range fields {
		keyVal := strings.Split(field, ":")
		if len(keyVal) != 2 {
			continue
		}
		data[keyVal[0]] = keyVal[1]
	}

	return data["ecl"] != "" && data["pid"] != "" && data["eyr"] != "" && data["hcl"] != "" && data["byr"] != "" && data["iyr"] != "" && data["hgt"] != ""
}
