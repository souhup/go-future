/*
给定一个非空二叉树，返回其最大路径和。

本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。

示例 1:

输入: [1,2,3]

       1
      / \
     2   3

输出: 6
示例 2:

输入: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

输出: 42
*/
package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: &TreeNode{Val: 1},
			},
		},
	}
	fmt.Println(maxPathSum(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	var globalMaxPathSum = -1 << 31
	myMaxPathSum(root, &globalMaxPathSum)
	return globalMaxPathSum
}

func myMaxPathSum(root *TreeNode, globalMaxPathSum *int) int {
	if root == nil {
		return 0
	}
	left := myMaxPathSum(root.Left, globalMaxPathSum)
	right := myMaxPathSum(root.Right, globalMaxPathSum)
	if left > 0 && right > 0 {
		*globalMaxPathSum = max(*globalMaxPathSum, root.Val+left+right)
		return root.Val + max(left, right)
	} else if left > 0 {
		*globalMaxPathSum = max(*globalMaxPathSum, root.Val+left)
		return root.Val + left
	} else if right > 0 {
		*globalMaxPathSum = max(*globalMaxPathSum, root.Val+right)
		return root.Val + right
	} else {
		*globalMaxPathSum = max(*globalMaxPathSum, root.Val)
		return root.Val
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
