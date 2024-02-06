package countcompletetreenodes

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func getLeftHeight(root *TreeNode) int {
	start := root 
	leftHeight := 0
	for start != nil {
		start = start.Left
		leftHeight++
	}

	return leftHeight
}

func getRightHeight(root *TreeNode) int {
	start := root 
	rightHeight := 0
	for start != nil {
		start = start.Right
		rightHeight++
	}

	return rightHeight
}

func countNodes(root *TreeNode) int {
	lh := getLeftHeight(root)
	rh := getRightHeight(root)
	if (lh == rh) {
		return (1 << lh) - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
