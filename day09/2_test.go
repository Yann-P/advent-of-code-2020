package main

import (
	"fmt"
	"testing"
)

func TestFindContiguousValuesThatSumTo(t *testing.T) {
	ans := FindContiguousValuesThatSumTo([]int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}, 127)
	fmt.Printf("%#v", ans)
}
