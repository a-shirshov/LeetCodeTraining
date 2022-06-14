package main

import (
	"twoSum/twoSum"
	"fmt"
)

func main() {
	fmt.Print("Enter length: ")
	var length int
	fmt.Scan(&length)

	fmt.Print("Enter Nums: ")
	nums := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Print("Enter target: ")
	var target int
	fmt.Scan(&target)

	result := twoSum.TwoSum(nums,target)
	fmt.Println("Result:",result)

	resultTwoPassHash := twoSum.TwoSum_twoPassHash(nums,target)
	fmt.Println("ResultTwoPassHash:",resultTwoPassHash)

	resultOnePassHash := twoSum.TwoSum_onePassHash(nums,target)
	fmt.Println("ResultOnePassHash:",resultOnePassHash)
}