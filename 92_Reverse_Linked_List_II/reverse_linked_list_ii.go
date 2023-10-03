package reverselinkedlistii

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
    dummy := &ListNode{
		Next: head,
	}
	previous := dummy
	current := head
	for i := 0; i < left - 1; i++ {
		previous, current = current, current.Next
	}

	leftPrevious := previous
	leftCurrent := current
	previous = nil
	for i := 0; i < right - left + 1; i++ {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	leftPrevious.Next = previous
	leftCurrent.Next = current
	return dummy.Next
}