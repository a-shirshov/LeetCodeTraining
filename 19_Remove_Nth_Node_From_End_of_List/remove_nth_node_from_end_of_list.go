package removenthnodefromendoflist

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{
		Next: head,
	}

	current := head
	count := 0
	for i := 0; i < n; i++ {
		current = current.Next
	}

}