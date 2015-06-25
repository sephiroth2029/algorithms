package threesum

import "testing"

func TestCount(t *testing.T) {
	var tsum ThreeSum
	sum := tsum.Count([]int{30, -40, -20, -10, 40, 0, 10, 5})
	if 4 != sum {
		t.Error("Count returned ", sum, ", expected ", 4)
	}
}
