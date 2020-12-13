package main

import "testing"

func TestGetCharacterRight3Down1(t *testing.T) {
	s := [...]string{
		"abcd",
		"efgh",
		"ijkl"}

	if GetCharacterRight3Down1(s[:], 0, 0) != 'h' {
		t.Error("Simple case")
	}

	if GetCharacterRight3Down1(s[:], 2, 1) != 'j' {
		t.Error("Wrap around")
	}
}
