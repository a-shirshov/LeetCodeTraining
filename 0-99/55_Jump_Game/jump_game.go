package jumpgame

func canJump(nums []int) bool {
	maxReach := 0
    for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}

		distance := i + nums[i]
		if distance > maxReach {
			maxReach = distance
		}
		
		if maxReach >= len(nums) - 1 {
			return true
		}
		
	}
	return true
}