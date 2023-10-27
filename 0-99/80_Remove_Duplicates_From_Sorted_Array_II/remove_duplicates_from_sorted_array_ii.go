package removeduplicatesfromsortedarrayii

func removeDuplicates(nums []int) int {
	amountOfSlide := 0

	currentElement := nums[0]
	amountOfReps := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == currentElement {
			if amountOfReps < 2 {
				amountOfReps++
				nums[i-amountOfSlide] = currentElement
			} else {
				amountOfSlide++
			}
		} else {
			currentElement = nums[i]
			nums[i-amountOfSlide] = currentElement
			amountOfReps = 1
		}
	}
	return len(nums) - amountOfSlide
}