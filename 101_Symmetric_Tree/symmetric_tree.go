package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}

    if root.Left == nil || root.Right == nil {
		return false
	}

	rightQueue := make([]*TreeNode, 0, 1)
	leftQueue := make([]*TreeNode, 0, 1)
	rightQueue = append(rightQueue, root.Right)
	leftQueue = append(leftQueue, root.Left)
	for len(leftQueue) != 0 && len(rightQueue) != 0 {
		leftPop := leftQueue[0]
		rightPop := rightQueue[0]
		leftQueue = leftQueue[1:]
		rightQueue = rightQueue[1:]

		if rightPop == nil && leftPop == nil {
			continue
		}

		if rightPop == nil || leftPop == nil {
			return false
		}

		if leftPop.Val != rightPop.Val {
			return false
		}

		if leftPop != nil {
			leftQueue = append(leftQueue, leftPop.Left)
			leftQueue = append(leftQueue, leftPop.Right)
		}

		if rightPop != nil {
			rightQueue = append(rightQueue, rightPop.Right)
			rightQueue = append(rightQueue, rightPop.Left)
		}
		
	}

	if len(leftQueue) == 0 && len(rightQueue) == 0 {
		return true
	}

	return false
}

func main() {
	isSymmetric(&TreeNode{
		Val: 1,
		Left: &TreeNode{2,&TreeNode{3, nil, nil},&TreeNode{4,nil,nil}},
		Right: &TreeNode{2,&TreeNode{4, nil, nil},&TreeNode{3,nil,nil}},
	})
}