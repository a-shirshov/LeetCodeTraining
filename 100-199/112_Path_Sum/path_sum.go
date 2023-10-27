package pathsum

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func addNode(root *TreeNode, targetSum int, currentResult int) bool {
	if root == nil {
		return false
	}
	currentResult += root.Val
	if currentResult == targetSum {
		if root.Left == nil && root.Right == nil {
			return true
		}
	}
	return addNode(root.Left, targetSum, currentResult) || addNode(root.Right, targetSum, currentResult)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	currentResult := 0
    
	return addNode(root, targetSum, currentResult)
}