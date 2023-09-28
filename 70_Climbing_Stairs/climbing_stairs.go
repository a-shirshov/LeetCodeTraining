package climbingstairs

func climbStairs(n int) int {
    if n <= 2 {
		return n
	}

	dpArray := []int{1,2}
	for count := 2; len(dpArray) < n; count++ {
		dpArray = append(dpArray, dpArray[count - 1] + dpArray[count - 2])
	}

	return dpArray[len(dpArray) - 1]
}