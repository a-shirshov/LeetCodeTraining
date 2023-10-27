package constructbinarytreefrompreorderandinordertraversal

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	var rootInorderIndex int
	for index, elem := range inorder {
		if preorder[0] == elem {
			rootInorderIndex = index
		}
	}
	root.Left = buildTree(preorder[1:rootInorderIndex + 1], inorder[:rootInorderIndex])
	root.Right = buildTree(preorder[rootInorderIndex + 1:], inorder[rootInorderIndex + 1:])
	return root
}