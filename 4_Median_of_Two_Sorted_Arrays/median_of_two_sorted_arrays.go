package median

import (
	"math"
)

func max(first, second int) int {
	if first > second {
		return first
	}
	return second
}

func min(first, second int) int {
	if first < second {
		return first
	}
	return second
}

func getLeftValue(index int, nums []int) int {
	if index < 0 {
		return math.MinInt
	}
	return nums[index]
}

func getRightValue(index int, nums []int) int {
	if index >= len(nums) {
		return math.MaxInt
	}
	return nums[index]
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	n1 := len(nums1)
	n2 := len(nums2)
	
	
	lPointer := 0
	rPointer := n1
	totalLength := n1 + n2
	partitionSize := (n1 + n2 + 1) / 2

	var l1, l2, r1, r2 int

	for {
		firstCut := (lPointer + rPointer) / 2
		
		secondCut := partitionSize - firstCut

		l1 = getLeftValue(firstCut - 1, nums1)
		r1 = getRightValue(firstCut, nums1)
		l2 = getLeftValue(secondCut - 1, nums2)
		r2 = getRightValue(secondCut, nums2)

		if (l1 > r2) {
			rPointer = firstCut - 1
			continue
		} else if (r1 < l2) {
			lPointer = firstCut + 1
			continue
		}

		break
	}
	if totalLength % 2 == 0 {
		return float64(min(r1,r2) + max(l1,l2)) / 2
	}
	
	return float64(max(l1,l2))
}
