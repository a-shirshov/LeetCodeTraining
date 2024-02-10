package validatebinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkValidBST(root *TreeNode, rootVal int) bool {
	if root == nil {
		return true
	}
	
	if root.Left != nil && root.Left.Val >= root.Val && root.Left.Val >= rootVal {
		return false
	}

	if root.Right != nil && root.Right.Val <= root.Val && root.Right.Val <= rootVal {
		return false
	}

	return checkValidBST(root.Left, rootVal) && checkValidBST(root.Right, rootVal)
}	

func isValidBST(root *TreeNode) bool {
	return checkValidBST(root, root.Val)
}