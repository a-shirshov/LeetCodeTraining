package main

import "fmt"

func jump(nums []int) int {
    steps := 0
	l, r := 0, 0
	
	for r < len(nums) - 1 {
		maxStep := 0
		for i := l; l <= r; i++ {
			if i + nums[i] > maxStep {
				maxStep = nums[i] + i
			}
			l++
		}
		steps++
		l = r + 1
		r = maxStep
	}
	return steps
}

func main() {
	fmt.Println(jump([]int{3,4,3,2,5,4,3}))
}