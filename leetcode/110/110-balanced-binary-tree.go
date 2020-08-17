/*
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。


*/
package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{Val: 15,
				Left: &TreeNode{
					Val: 16,
				}},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(isBalanced(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return treeDepth(root) != -1
}

func treeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := treeDepth(root.Left)
	rightDepth := treeDepth(root.Right)
	if leftDepth == -1 || rightDepth == -1 {
		return -1
	}
	if rightDepth-leftDepth > 1 || leftDepth-rightDepth > 1 {
		return -1
	}
	return 1 + max(leftDepth, rightDepth)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
