package all

// Problem: 701
// Title: Insert into a Binary Search Tree
// Category: all
// Tags: all


import "github.com/halfrost/LeetCode-Go/structures"

// TreeNode define
type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insert(n *TreeNode, val int) *TreeNode {
	if n == nil {
		return &TreeNode{Val: val}
	}
	if n.Val < val {
		n.Right = insert(n.Right, val)
	} else {
		n.Left = insert(n.Left, val)
	}
	return n
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	return insert(root, val)
}
