package subtreeofanothertree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}	

func sameTree(firstRoot *TreeNode, secondRoot *TreeNode) bool {
	if firstRoot == nil && secondRoot == nil {
		return true
	}
	
	if firstRoot != nil && secondRoot != nil && firstRoot.Val == secondRoot.Val {
		return sameTree(firstRoot.Left, secondRoot.Left) && 
			sameTree(firstRoot.Right, secondRoot.Right)
	}

	return false
	
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if subRoot == nil {
		return true
	}
	
	if root == nil {
		return false
	}

	if sameTree(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}