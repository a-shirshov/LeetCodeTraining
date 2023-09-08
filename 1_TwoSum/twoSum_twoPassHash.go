package twoSum

func TwoSum_twoPassHash(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hashTable[nums[i]] = i;
	}

	var result []int

	for i := 0; i < len(nums); i++ {
		brother_number := target - nums[i];
		brother_index, ok := hashTable[brother_number]
		if ok && (brother_index != i) {
			result = append(result, i)
			result = append(result, brother_index)
			break
		}
	}

	return result
}