package minimumsizesubarraysum

func minSubArrayLen(target int, nums []int) int {
	left := 0
	right := 0
	currentSum := 0
	minLength := len(nums) + 1
    for right < len(nums){
		currentSum += nums[right]
		if currentSum >= target {
			for currentSum >= target && left <= right {
				currentLength := right - left + 1
				if currentLength < minLength {
					minLength = currentLength
				}

				currentSum -=nums[left]
				left++
			}
		}
		right++
	}

	if minLength > len(nums) {
		return 0
	}

	return minLength
}