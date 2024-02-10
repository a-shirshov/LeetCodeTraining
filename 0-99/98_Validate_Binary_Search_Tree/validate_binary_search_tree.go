package main

import "math"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func checkValid(root *TreeNode, left, right int) bool {
	if root == nil {
		return true
	}

	if root.Val <= left || root.Val >= right {
		return false
	} 

	return checkValid(root.Left, left, root.Val) && checkValid(root.Right, root.Val, right)
}

func isValidBST(root *TreeNode) bool {
	return checkValid(root, math.MinInt, math.MaxInt )
}