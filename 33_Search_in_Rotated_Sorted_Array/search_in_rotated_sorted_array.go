package searchinrotatedsortedarray

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		}

        //left sorted part
		if nums[m] >= nums[0] {
            if (target >= nums[l] && target < nums[m]) {
                r = m - 1
            } else {
                l = m + 1
            }
        } else {
            if (target <= nums[r] && target > nums[m]) {
                l = m + 1
            } else {
                r = m - 1
            }
        }
	}

	return -1
}
