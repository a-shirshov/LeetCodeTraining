package containsduplicateii

func containsNearbyDuplicate(nums []int, k int) bool {
    hashMap := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		_, ok := hashMap[nums[i]]
		if ok {
			return true
		}
		hashMap[nums[i]] = struct{}{}
		if len(hashMap) > k {
			delete(hashMap, nums[i-k])
		}
	}
	return false
}