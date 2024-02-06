package binarytreelevelordertraversal

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{}
	result := [][]int{}
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		resLine := []int{}
		for i := 0; i < size; i++ {
			elem := queue[0]
			queue = queue[1:]
			if elem == nil {
				continue
			}
			resLine = append(resLine, elem.Val)
			queue = append(queue, elem.Left)
			queue = append(queue, elem.Right)
		}
		if len(resLine) != 0 {
			result = append(result, resLine)
		}
	}
	return result
}