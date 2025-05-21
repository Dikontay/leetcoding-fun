package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	piles := []int{30, 11, 23, 4, 20}
	h := 5
	fmt.Println(minEatingSpeed(piles, h))
}

// 875.Koko is eating bananas
func minEatingSpeed(piles []int, h int) int {

	m := slices.Max(piles)

	left := 1
	right := m
	var mid int
	var currEatingHours int
	var k int
	var minK int
	minK = m
	for left <= right {
		mid = (left + right) / 2

		for i := range piles {
			currEatingHours += int(math.Round(float64(piles[i]) / float64(mid)))
		}
		if int(currEatingHours) <= h {
			k = mid
			right = mid - 1

		} else {
			left = mid + 1
		}
		if minK > k {
			minK = k
		}
		currEatingHours = 0
	}
	return k

}
