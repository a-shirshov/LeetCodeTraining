package binarytreerightsideview

type TreeNode struct {
 	Val int
	Left *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	result := []int{}
	queue := []*TreeNode{}
	lc := -1
	queue = append(queue, root)
	countQueue := []int{0}
	for len(queue) != 0 {
		elem := queue[0]
		queue = queue[1:]
		countElem := countQueue[0]
		countQueue = countQueue[1:]
		if elem == nil {
			continue
		}

		if countElem != lc {
			result = append(result, elem.Val)
			lc = countElem
		}

		queue = append(queue, elem.Right)
		countQueue = append(countQueue, countElem+1)
		queue = append(queue, elem.Left)
		countQueue = append(countQueue, countElem+1)
	}
	return result
}