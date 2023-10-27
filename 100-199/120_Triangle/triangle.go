package triangle

func min(first, second int) int {
	if first < second {
		return first
	}
	return second
}

func minimumTotal(triangle [][]int) int {
    if len(triangle) == 0 {
		return 0
	}

	if len(triangle) == 1 {
		return triangle[0][0]
	}

	dpArray := make([]int,1,len(triangle))
	dpArray[0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		dpArray = append(dpArray, 100000) //>10^4
	}

	var currentMin int
	for i := 1; i < len(triangle); i++ {
		currentMin = 100000
		for j := len(triangle[i]) - 1; j >=0; j-- {
			if j != 0 {
				dpArray[j] = min(dpArray[j], dpArray[j-1])
			}
			dpArray[j] += triangle[i][j]
			currentMin = min(currentMin, dpArray[j])
		}
	}
	return currentMin
}