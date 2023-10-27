package rotatearray

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func rotate(nums []int, k int)  {
	k = k % len(nums)
    reverse(nums, 0, len(nums) - 1)
	reverse(nums, 0, k - 1)
	reverse(nums, k, len(nums) - 1)
}