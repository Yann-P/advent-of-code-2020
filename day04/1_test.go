package main

import "testing"

func TestIsPassportValid(t *testing.T) {
	if IsPassportValid("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm") != true {
		t.Error("Passport 1 should be valid")
	}
	if IsPassportValid("iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929") != false {
		t.Error("Passport 2 should be invalid")
	}
	if IsPassportValid("hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm") != true {
		t.Error("Passport 3 should be valid")
	}
	if IsPassportValid("hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in") != false {
		t.Error("Passport 4 should be invalid")
	}
}
