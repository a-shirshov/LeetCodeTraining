package rotatelist

type ListNode struct {
	Val int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
		return head
	}
	current := head
	previous := head
	length := 0
	for current != nil {
		length++
		previous = current
		current = current.Next
	}
	k = k % length
	if k == 0 {
		return head
	}
	
	previous.Next = head
	current = head
	for i := 0; i < length - k-1; i++ {
		current = current.Next
	}
	newHead := current.Next
	current.Next = nil
	return newHead
}