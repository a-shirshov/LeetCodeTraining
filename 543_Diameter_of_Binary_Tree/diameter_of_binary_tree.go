package diameterofbinarytree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func max(first, second int) int {
	if first >= second {
		return first
	}
	return second
}

func dfs(root *TreeNode, maxLength *int) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left, maxLength)
	right := dfs(root.Right, maxLength)
	*maxLength = max(*maxLength,  left + right)

	return max(left, right) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxLength := 0
	_ = dfs(root, &maxLength)
	return maxLength
}