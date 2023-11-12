package summaryranges

import "fmt"

func summaryRanges(nums []int) []string {
    var result []string
	if len(nums) == 0 {
		return result
	}
	start, previous := nums[0], nums[0] 
	for i := 1; i < len(nums); i++ {
		if nums[i] - previous == 1 {
			previous = nums[i]
			continue
		}

		if start == previous {
			result = append(result, fmt.Sprintf("%d", start))
		} else {
			result = append(result, fmt.Sprintf("%d->%d", start, previous))
		}
		start, previous = nums[i], nums[i]
	}

	if start == previous {
		result = append(result, fmt.Sprintf("%d", start))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", start, previous))
	}

	return result
}