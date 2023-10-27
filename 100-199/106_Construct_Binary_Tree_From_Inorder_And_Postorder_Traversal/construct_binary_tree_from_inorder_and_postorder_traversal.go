package constructbinarytreefrominorderandpostordertraversal

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 && len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: postorder[len(postorder) - 1],
	}
	var rootInorderIndex int
	for index, elem := range inorder {
		if elem == postorder[len(postorder) - 1] {
			rootInorderIndex = index
		}
	}
	root.Left = buildTree(inorder[:rootInorderIndex], postorder[:rootInorderIndex])
	root.Right = buildTree(inorder[rootInorderIndex + 1:], postorder[rootInorderIndex:len(postorder) - 1])
	return root
}