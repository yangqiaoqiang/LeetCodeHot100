package main

import "container/list"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	l := list.New()
	l.PushBack(root)
	res := [][]int{}
	kk := []int{}
	for l.Len() > 0 {
		sum := l.Len()
		for i := 0; i < sum; i++ {
			node := l.Remove(l.Front()).(*TreeNode)
			if node.Left != nil {
				l.PushBack(node.Left)
			}
			if node.Right != nil {
				l.PushBack(node.Right)
			}
			kk = append(kk, node.Val)
		}
		res = append(res, kk)
		kk = []int{}
	}
	return res
}
