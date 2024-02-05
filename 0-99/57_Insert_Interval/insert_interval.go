package insertinterval

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
    if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	var result [][]int
	previous := intervals[0]
	hasUsedNew := false
	if newInterval[1] < previous[0] {
		result = append(result, newInterval)
		hasUsedNew = true
	}
	
	for i := 1; i < len(intervals); i++ {
		if !hasUsedNew && newInterval[0] <= previous[1]{
			if newInterval[1] < previous[0]{
				result = append(result, newInterval)
			} else {
				previous[0] = min(previous[0], newInterval[0])
				previous[1] = max(previous[1], newInterval[1])
			}

			hasUsedNew = true
		}

		if intervals[i][0] <= previous[1] {
			intervals[i][0] = min(intervals[i][0], previous[0])
			intervals[i][1] = max(intervals[i][1], previous[1])
		} else {
			result = append(result, previous)
		}
		previous = intervals[i]
	}

	if !hasUsedNew && newInterval[0] <= previous[1]{
		if newInterval[1] < previous[0]{
			result = append(result, newInterval)
		} else {
			previous[0] = min(previous[0], newInterval[0])
			previous[1] = max(previous[1], newInterval[1])
		}
		hasUsedNew = true
	}
	
	result = append(result, previous)
	if !hasUsedNew {
		result = append(result, newInterval)
	}
	return result
}