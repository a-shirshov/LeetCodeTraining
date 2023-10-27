package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	result := make([][]int, 0, 64)
	result = append(result, []int{})
	for i := 0; i < len(nums); i++ {
		length := len(result)
		temp := make([][]int, length)
		for j := 0; j < length; j++ {
			sliceToAppend := append(result[j], nums[i])
			temp[j] = sliceToAppend
			fmt.Printf("res after = %v\n\n", result)
		}
		result = append(result, temp...)
	}
	return result
}

func main() {
	test := []int{1, 2, 3, 4, 5}
	subsets(test)
}