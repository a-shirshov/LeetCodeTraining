package mergeintervals

import "sort"

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
} 

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {return intervals[i][0] < intervals[j][0]})
	var result [][]int
	previous := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= previous[0] && intervals[i][0] <= previous[1] {
			intervals[i] = []int{previous[0], max(previous[1], intervals[i][1])}
		} else {
			result = append(result, previous)
		}
		previous = intervals[i]
	}
	result = append(result, previous)
	return result
}