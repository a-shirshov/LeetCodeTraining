package topkfrequentelements

func topKFrequent(nums []int, k int) []int {
	hashMap := make(map[int]int)
	countSlice := make([][]int, len(nums) + 1)

	for i := 0; i<len(nums); i++{
		if _, ok := hashMap[nums[i]]; !ok {
			hashMap[nums[i]] = 0
		}
		hashMap[nums[i]]++
	}
	for number, count := range hashMap {
		countSlice[count] = append(countSlice[count], number)
	}
	
	var result []int
	for i := len(countSlice)-1; i > 0; i-- {
		result = append(result, countSlice[i]...)
		if len(result) == k {
			return result
		}
	}
	return result
}