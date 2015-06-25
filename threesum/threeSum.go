package threesum

import "fmt"

// ThreeSum implements the current algorithm for the 3-sum problem.
type ThreeSum struct {
}

// Count returns the number of triplets that sum zero.
func (tsum *ThreeSum) Count(a []int) int {
	N := len(a)
	count := 0
	iterations := 0

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := j + 1; k < N; k++ {
				iterations++
				if a[i]+a[j]+a[k] == 0 {
					count++
				}
			}
		}
	}

	fmt.Println(iterations)

	return count
}
