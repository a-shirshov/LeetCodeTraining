package removenthnodefromendoflist

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{
		Next: head,
	}

	movePointer := dummy
	for i := 0; i < n; i++ {
		movePointer = movePointer.Next
	}

	current := dummy
	for movePointer.Next != nil {
		movePointer = movePointer.Next
		current = current.Next
	}
	current.Next = current.Next.Next
	return dummy.Next
}