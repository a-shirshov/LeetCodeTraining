package copylistwithrandompointer

type Node struct {
	Val int
	Next *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
    oldToCopy := make(map[*Node]*Node)
	current := head
	for current != nil {
		copy := &Node{
			Val: current.Val,
		}
		oldToCopy[current] = copy
		current = current.Next
	}

	current = head
	for current != nil {
		copy := oldToCopy[current]
		copy.Next = oldToCopy[current.Next]
		copy.Random = oldToCopy[current.Random]
		current = current.Next
	}

	return oldToCopy[head]
}