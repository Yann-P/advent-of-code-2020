package main

import "testing"

func TestParseRowNumber(t *testing.T) {
	got := ParseRowNumber("FBFBBFF")
	expect := 44
	if got != expect {
		t.Errorf("expected %d; got %d", expect, got)
	}
}

func TestParseColNumber(t *testing.T) {
	got := ParseColNumber("RLR")
	expect := 5
	if got != expect {
		t.Errorf("expected %d; got %d", expect, got)
	}
}

func TestExtractSeatId(t *testing.T) {
	got := ExtractSeatId("BFFFBBFRRR")
	expect := 567
	if got != expect {
		t.Errorf("expected %d; got %d", expect, got)
	}

	got = ExtractSeatId("FFFBBBFRRR")
	expect = 119
	if got != expect {
		t.Errorf("expected %d; got %d", expect, got)
	}

	got = ExtractSeatId("BBFFBBFRLL")
	expect = 820
	if got != expect {
		t.Errorf("expected %d; got %d", expect, got)
	}
}

func TestExtractRowNumberFromSeatId(t *testing.T) {
	got := ExtractRowNumberFromSeatId(ExtractSeatId("BBFFBBFRLL"))
	expect := 102
	if got != expect {
		t.Errorf("expected %d; got %d", expect, got)
	}
}
