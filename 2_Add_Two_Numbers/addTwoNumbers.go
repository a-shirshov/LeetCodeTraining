package addTwoNumbers

import (
	"fmt"
)

func PrintList(list *ListNode) {
	for list != nil {
		fmt.Print(list.Val)
		list = list.Next
	}
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 0
	saved_digits := 0
	var list *ListNode

	head := list
	
	for (l1 != nil) && (l2 != nil) {
		sum = l1.Val + l2.Val
		number := (sum + saved_digits) % 10

		Node := &ListNode{
			number,
			nil,
		}

		l1 = l1.Next
		l2 = l2.Next

		if head == nil {
			head = Node
			list = head 
		} else {
			list.Next = Node
			list = list.Next
		}

		saved_digits = (sum + saved_digits) / 10
	}
	
	for (l1 != nil) {
		number := (l1.Val + saved_digits) % 10

		Node := &ListNode{
			number,
			nil,
		}
		saved_digits = (l1.Val + saved_digits) / 10

		list.Next = Node
		l1 = l1.Next
		list = list.Next
	}

	for (l2 != nil) {
		number := (l2.Val + saved_digits) % 10

		Node := &ListNode{
			number,
			nil,
		}

		saved_digits = (l2.Val + saved_digits) / 10

		list.Next = Node
		l2 = l2.Next
		list = list.Next
	}

	if saved_digits != 0 {
		Node := &ListNode{
			saved_digits,
			nil,
		}
		list.Next = Node
	}

	return head
}