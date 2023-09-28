package balancedbinarytree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func abs(first, second int) (int) {
	res := first-second
	if res < 0 {
		return -res
	}

	return res
}

func max(first, second int) int {
	if first >= second {
		return first
	}
	return second
}

func dfs(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	leftIsBalanced, leftHeight := dfs(root.Left)
	rightIsBalanced, rightHeight := dfs(root.Right)
	balanced := (leftIsBalanced && rightIsBalanced) && abs(leftHeight, rightHeight) <= 1 
	return balanced, 1 + max(leftHeight, rightHeight)
}


func isBalanced(root *TreeNode) bool {
    balanced, _ := dfs(root)
	return balanced
}