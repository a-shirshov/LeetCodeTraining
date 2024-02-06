package averageoflevelsinbinarytree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
    queue := []*TreeNode{}
	queue = append(queue, root)
	result := []float64{}
	for len(queue) != 0 {
		size := len(queue)
		sum := 0
		count := 0
		for i := 0; i < size; i++ {
			elem := queue[0]
			queue = queue[1:]
			if elem == nil {
				continue
			}
			sum += elem.Val
			count++
			queue = append(queue, elem.Left)
			queue = append(queue, elem.Right)
		}
		if count != 0 {
			result = append(result, float64(sum)/float64(count))
		}
	}
	return result
}