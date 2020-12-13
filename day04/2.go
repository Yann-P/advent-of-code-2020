package main

import (
	"regexp"
	"strconv"
	"strings"
)

func IsFourDigitNumbersBetween(s string, minIncl int, maxInc int) bool {
	if len(s) != 4 {
		return false
	}

	num, err := strconv.Atoi(s)

	if err != nil {
		return false
	}

	return num >= minIncl && num <= maxInc
}

func IsByrValid(s string) bool {
	return IsFourDigitNumbersBetween(s, 1920, 2002)
}

func IsIyrValid(s string) bool {
	return IsFourDigitNumbersBetween(s, 2010, 2020)
}

func IsEyrValid(s string) bool {
	return IsFourDigitNumbersBetween(s, 2020, 2030)
}

func IsHgtValid(s string) bool {
	r, err := regexp.Compile("^(\\d+)(cm|in)$")
	if err != nil {
		panic(err)
	}

	extracted := r.FindStringSubmatch(s) // [180cm, 180, cm]

	if len(extracted) != 3 {
		return false
	}

	heightStr := extracted[1]
	cmOrIn := extracted[2]
	height, err := strconv.Atoi(heightStr)

	if err != nil {
		panic(err)
	}

	return (cmOrIn == "cm" && height >= 150 && height <= 193) || (cmOrIn == "in" && height >= 59 && height <= 76)
}

func IsHclValid(s string) bool {
	r, err := regexp.Compile("^#[0-9a-f]{6}$")
	if err != nil {
		panic(err)
	}

	return r.MatchString(s)
}

func IsEclValid(s string) bool {
	return s == "amb" || s == "blu" || s == "brn" || s == "gry" || s == "grn" || s == "hzl" || s == "oth"
}

func IsPidValid(s string) bool {
	r, err := regexp.Compile("^[0-9]{9}$")
	if err != nil {
		panic(err)
	}

	return r.MatchString(s)
}

func IsPassportValidStrict(passport string) bool {

	fields := strings.Split(passport, " ")

	data := make(map[string]string)

	for _, field := range fields {
		keyVal := strings.Split(field, ":")
		if len(keyVal) != 2 {
			continue
		}
		data[keyVal[0]] = keyVal[1]
	}

	return data["ecl"] != "" && IsEclValid(data["ecl"]) &&
		data["pid"] != "" && IsPidValid(data["pid"]) &&
		data["eyr"] != "" && IsEyrValid(data["eyr"]) &&
		data["hcl"] != "" && IsHclValid(data["hcl"]) &&
		data["byr"] != "" && IsByrValid(data["byr"]) &&
		data["iyr"] != "" && IsIyrValid(data["iyr"]) &&
		data["hgt"] != "" && IsHgtValid(data["hgt"])
}
