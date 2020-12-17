package main

import "testing"

func TestIsNumberSumOfTwo(t *testing.T) {
	if IsNumberSumOfTwo([]int{1, 2, 3}, 4) != true {
		t.Error(1)
	}

	if IsNumberSumOfTwo([]int{1, 2, 3}, 7) != false {
		t.Error(2)
	}
}

func TestFindFirstNumberNotSumOf(t *testing.T) {
	ans := FindFirstNumberNotSumOf([]int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}, 5)

	if ans != 127 {
		t.Error(1, ans)
	}
}
