package sametree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var queueP []*TreeNode
	var queueQ []*TreeNode
	queueP = append(queueP, p)
	queueQ = append(queueQ, q)

	for len(queueP) != 0 && len(queueQ) != 0 {
		nodeP := queueP[0]
		nodeQ := queueQ[0]
		
		queueP = queueP[1:]
		queueQ = queueQ[1:]

		if nodeP == nil && nodeQ == nil {
			continue
		}

		if nodeP == nil || nodeQ == nil {	
			return false
		}

		if nodeP.Val != nodeQ.Val {
			return false
		}

		queueP = append(queueP, nodeP.Left)
		queueQ = append(queueQ, nodeQ.Left)

		queueP = append(queueP, nodeP.Right)
		queueQ = append(queueQ, nodeQ.Right)
	}

	if len(queueP) != 0 || len(queueQ) != 0 {
		return false
	}

	return true
}