package twosumiiinputarrayissorted

func twoSum(numbers []int, target int) []int {
    firstPointer := 0
	secondPointer := len(numbers) - 1
	for firstPointer < secondPointer {
		sum := numbers[firstPointer] + numbers[secondPointer]
		
		if sum > target {
			secondPointer--
			continue
		}

		if sum < target {
			firstPointer++
			continue
		}

		return []int{firstPointer+1, secondPointer+1}
	}
	return []int{}
}