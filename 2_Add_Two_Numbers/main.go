package main

import (
	"addTwoNumbers/addTwoNumbers"
)

func main() {
	result := addTwoNumbers.AddTwoNumbers(
		&addTwoNumbers.ListNode{2, &addTwoNumbers.ListNode{4, &addTwoNumbers.ListNode{3, nil}}},
		&addTwoNumbers.ListNode{5, &addTwoNumbers.ListNode{6, &addTwoNumbers.ListNode{4, nil}}}, 
	)
	addTwoNumbers.PrintList(result)
}