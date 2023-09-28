package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	currentElement := -101
	amountOfReps := 0
    for i := 0; i < len(nums); i++ {
		if nums[i] != currentElement {
			currentElement = nums[i]
			nums[i-amountOfReps] = currentElement
			continue
		}
		amountOfReps++
	}
	return len(nums) - amountOfReps
}