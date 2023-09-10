package maximumdepthofbinarytree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func maxDepth(root *TreeNode) int {
	length := 0
	if root != nil {
		length = maxDepth(root.Left)
		length = maxDepth(root.Right)
	}
   
	return length + 1
}
