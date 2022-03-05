package main

import "container/list"

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := list.New()
	l.PushBack(root)
	res := 0
	for l.Len() > 0 {
		sum := l.Len()
		res++
		for i := 0; i < sum; i++ {
			node := l.Remove(l.Front()).(*TreeNode)
			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
		}
	}
	return res
}
