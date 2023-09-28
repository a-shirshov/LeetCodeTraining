package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	saveRoot := root
	if root != nil {
		_ = invertTree(root.Left)
		_ = invertTree(root.Right)
		root.Left, root.Right = root.Right, root.Left
	}
	return saveRoot
}

func main() {
	tree := &TreeNode{
		4, 
		&TreeNode{
			2, 
			&TreeNode{1, nil, nil},
			&TreeNode{3, nil, nil},
		}, 
		&TreeNode{
			7, 
			&TreeNode{6, nil, nil}, 
			&TreeNode{9, nil, nil},
		},
	}
	_ = invertTree(tree)
}