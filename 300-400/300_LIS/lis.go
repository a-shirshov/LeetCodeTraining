package lis

func lengthOfLIS(nums []int) int {
    lis := make([]int, len(nums))
	max := 0
	for i := len(nums) - 1; i >= 0; i-- {
		lis[i] = 1
		for j := i+1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				if lis[j] + 1 > lis[i] {
					lis[i] = lis[j] + 1
				}
			}
		}
		if lis[i] > max {
			max = lis[i]
		}
	}
	return max
}