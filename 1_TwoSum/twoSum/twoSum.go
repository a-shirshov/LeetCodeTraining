package twoSum

func TwoSum(nums []int, target int) []int {
	var result []int
	for i := 0; i < len(nums)-1; i++ {
		first := nums[i]
		for j := i + 1; j < len(nums); j++ {
			second := nums[j]
			if second + first == target {
				result = append(result, i)
				result = append(result, j)
				return result
			}
		}
	}
	return result
}
