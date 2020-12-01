package main

import "testing"

func TestEmptyArray(t *testing.T) {
	arr := [...]int{}
	ans := MultiplyNumbersThatSumTo2020(arr[:])
	if ans != -1 {
		t.Errorf("findNumbersThatSumTo2020(%v) = %d; want %d", arr, ans, -1)
	}
}

func TestSingleValueArray(t *testing.T) {
	arr := [...]int{1}
	ans := MultiplyNumbersThatSumTo2020(arr[:])
	if ans != -1 {
		t.Errorf("findNumbersThatSumTo2020(%v) = %d; want %d", arr, ans, -1)
	}
}

func TestMultipleValueArray(t *testing.T) {
	arr := [...]int{2, 1010, 3, 1010}
	ans := MultiplyNumbersThatSumTo2020(arr[:])
	expect := 1010 * 1010
	if ans != expect {
		t.Errorf("findNumbersThatSumTo2020(%v) = %d; want %d", arr, ans, expect)
	}
}
