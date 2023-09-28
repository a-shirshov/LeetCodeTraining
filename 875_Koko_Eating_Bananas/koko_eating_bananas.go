package main

import (
	"math"
)
  
func minEatingSpeed(piles []int, h int) int {
	r := 0
	for _, pile := range piles {
		if (pile > r) {
			r = pile
		}
	}

	result := r
	l := 1
	for l<=r {
		current_hours := 0
		m := (l+r)/2
		for _, pile := range piles {
			current_hours += int(math.Ceil(float64(pile)/float64(m)))
		}
		
		if current_hours > h {
			l = m + 1
		} else {
			r = m - 1
			if m < result {
				result = m
			}
		}

	}

	return result
}

func main() {
	println(minEatingSpeed([]int{3,6,7,11}, 8))
}