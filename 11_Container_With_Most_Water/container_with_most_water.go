package containerwithmostwater

func min(first, second int) int {
	if second > first {
		return first
	}
	return second
}

func maxArea(height []int) int {
	firstPointer := 0
	secondPointer := len(height) - 1
	maxAmount := 0
	for firstPointer < secondPointer {
		length := secondPointer - firstPointer
		minHeight := min(height[firstPointer], height[secondPointer])
		waterAmount := minHeight * length
		if waterAmount > maxAmount {
			maxAmount = waterAmount
		}

		if (height[secondPointer] < height[firstPointer]){
			secondPointer-- 
			continue
		}

		firstPointer++
	}
	return maxAmount
}