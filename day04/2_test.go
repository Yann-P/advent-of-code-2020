package main

import "testing"

func TestIsHgtValid(t *testing.T) {
	if IsHgtValid("180cm") != true {
		t.Error("1 should be valid")
	}
}

func TestIsHclValid(t *testing.T) {
	if IsHclValid("#abcdef") != true {
		t.Error("1 should be valid")
	}
	if IsHclValid("#azcdef") != false {
		t.Error("2 should be invalid")
	}
}

func TestIsPidValid(t *testing.T) {
	if IsPidValid("000000001") != true {
		t.Error("1 should be valid")
	}
	if IsPidValid("0123456789") != false {
		t.Error("2 should be invalid")
	}
}
