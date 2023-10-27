package hindex

import "sort"

func hIndex(citations []int) int {
    sort.IntSlice(citations).Sort()

	hIndex := 0
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] >= hIndex + 1 {
			hIndex++
		} else {
			break
		}
	}
	return hIndex
}