package main

type ListNode struct {
     Val int
     Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy :=  &ListNode{};
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 == nil {
		current.Next = list2
	} 

	if list2 == nil {
		current.Next = list1
	}

	return dummy.Next
}

func main() {
	list1 := &ListNode{1, &ListNode{3, &ListNode{5, nil}}}
	list2 := &ListNode{2, &ListNode{4, &ListNode{6, nil}}}
	newList := mergeTwoLists(list1,list2)
	for newList != nil {
		println(newList.Val)
		newList = newList.Next
	}
}