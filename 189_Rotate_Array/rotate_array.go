package rotatearray

func rotate(nums []int, k int)  {
    for i := 0; i+len(nums)-k < len(nums); i++ {
		nums[i], nums[i+len(nums)-k] = nums[i+len(nums)-k], nums[i]
	}
}