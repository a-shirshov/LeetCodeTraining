package sum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	result := [][]int{}
	if length < 3 {
		return result
	}

	for firstNumberIndex := 0; firstNumberIndex < length- 2; firstNumberIndex++{
		firstNumber := nums[firstNumberIndex]
		if firstNumber > 0 {
			return result
		}

		firstPointer := firstNumberIndex + 1
		if firstNumberIndex != 0 && firstNumber == nums[firstNumberIndex - 1] {
			continue
		}

		secondPointer := length - 1 
		for firstPointer < secondPointer {
			sum := nums[firstPointer] + nums[secondPointer] + firstNumber
			if (sum > 0) {
				secondPointer--
				continue
			}

			if (sum < 0) {
				firstPointer++
				continue
			}
			
			result = append(result, []int{firstNumber, nums[firstPointer], nums[secondPointer]})
			secondPointer--
			for firstPointer < secondPointer && nums[secondPointer] == nums[secondPointer+1] {
				secondPointer--
			}
		}

	}
	return result
}