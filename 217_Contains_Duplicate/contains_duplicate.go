package containsduplicate

func containsDuplicate(nums []int) bool {
    hashMap := make(map[int]bool)
	for _, element := range nums {
		if hashMap[element] {
			return true
		}
		hashMap[element] = true
	}
	return false
}