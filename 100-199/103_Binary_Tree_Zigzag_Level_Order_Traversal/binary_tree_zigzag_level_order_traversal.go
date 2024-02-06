package binarytreezigzaglevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{}
	result := [][]int{}
	queue = append(queue, root)
	level := 1
	for len(queue) != 0 {
		size := len(queue)
		resLine := []int{}
		for i := 0; i < size; i++ {
			elem := queue[0]
			queue = queue[1:]
			if elem == nil {
				continue
			}
			resLine = append(resLine, elem.Val)
			queue = append(queue, elem.Left)
			queue = append(queue, elem.Right)

		}
		if len(resLine) != 0 {
			if level%2 == 0 {
				for i := 0; i < len(resLine)/2; i++ {
					resLine[i], resLine[len(resLine)-1-i] = resLine[len(resLine)-1-i], resLine[i]
				}
			}
			result = append(result, resLine)
		}
		level++
	}
	return result
}
