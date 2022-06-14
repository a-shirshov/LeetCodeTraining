package twoSum


func TwoSum_onePassHash(nums []int, target int) []int {
	hashTable := make(map[int]int)
	var result []int

	for i := 0; i < len(nums); i++ {
		brother_number := target - nums[i];
		brother_index, ok := hashTable[brother_number]
		if ok && (brother_index != i) {
			result = append(result, i)
			result = append(result, brother_index)
			break
		}
		hashTable[nums[i]] = i;
	}
	return result;
}