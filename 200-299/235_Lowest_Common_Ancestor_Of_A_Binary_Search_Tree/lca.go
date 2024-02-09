package lowestcommonancestorofabinarysearchtree



type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getPath(root, node *TreeNode) []*TreeNode {
	current := root
	var path []*TreeNode
	for current != node {
		path = append(path, current)
		if node.Val < current.Val {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	path = append(path, node)
	return path
}


// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	firstPath := getPath(root, p)
// 	secondPath := getPath(root, q)
// 	index := 0
// 	lca := root
// 	for index < len(firstPath) && index < len(secondPath) {
// 		if firstPath[index] == secondPath[index] {
// 			lca = firstPath[index]
// 		}
// 		index++
// 	}

// 	return lca
// }

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	current := root
	for current != nil {
		if p.Val < current.Val && q.Val < current.Val {
			current = current.Left
		} else if p.Val > current.Val && q.Val > current.Val {
			current = current.Right
		} else {
			return current
		}
	}
	return current
}