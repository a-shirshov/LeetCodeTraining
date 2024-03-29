package main

import "fmt"

func productExceptSelf(nums []int) []int {
    result := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := len(nums) - 1; i >=0; i-- {
		result[i] *= postfix
		postfix *= nums[i]
	}

	return result
}

func main(){
	nums := []int{1,2,3,4}
	fmt.Println(productExceptSelf(nums))
}