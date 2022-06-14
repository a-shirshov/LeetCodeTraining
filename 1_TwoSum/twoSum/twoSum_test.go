package twoSum

import (
	"testing"
)

var twoSumTests = []struct {
	nums   []int
	target int
	result []int
}{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{3, 2, 4}, 6, []int{1, 2}},
	{[]int{3, 3}, 6, []int{0, 1}},
}

func TestTwoSum(t *testing.T) {
	for _, test := range twoSumTests {
		t.Run("TwoSum", func(t *testing.T) {
			result := TwoSum(test.nums, test.target)

			if len(result) != len(test.result) {
				t.Errorf("Length problem")
			}

			for index := range result {
				if result[index] != test.result[index] {
					t.Errorf("got %v, want %v", result[index], test.result[index])
					break
				}
			}

		})
	}
}
