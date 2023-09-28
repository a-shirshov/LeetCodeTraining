package maximumdepthofbinarytree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func max(first, second int) int {
	if first >= second {
		return first
	}
	return second
}

func maxDepthRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	
	return 1 + max(maxDepthRecursive(root.Left), maxDepthRecursive(root.Right))
}

func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var queue []*TreeNode

	level := 0
	queue = append(queue, root)
	for len(queue) != 0 {
		for _, node := range queue {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			queue = queue[1:]
		}

		level++
	}

	return level
}

func maxDepthIDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type TreeNodeWithDepth struct {
		*TreeNode
		depth int
	}

	var stack []*TreeNodeWithDepth
	maxdepth := 0

	stack = append(stack, &TreeNodeWithDepth{root, 1})

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if maxdepth < node.depth {
			maxdepth = node.depth
		}

		if node.Right != nil {
			stack = append(stack, &TreeNodeWithDepth{node.Right, node.depth + 1})
		}

		if node.Left != nil {
			stack = append(stack, &TreeNodeWithDepth{node.Left, node.depth + 1})
		}
	}

	return maxdepth
}
