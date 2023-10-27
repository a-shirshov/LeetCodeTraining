package removeelement

func removeElement(nums []int, val int) int {
    amountOfReps := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			amountOfReps++
			continue
		}
		nums[i-amountOfReps] = nums[i]
	}
	return len(nums) - amountOfReps
}