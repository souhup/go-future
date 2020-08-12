/*
给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例:

输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
解释:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
*/
package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}},
	}
	fmt.Println(rightSideView(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var result = make([]int, 0)
	dfs(root, 0, &result)
	return result
}

func dfs(root *TreeNode, layer int, result *[]int) {
	if root == nil {
		return
	}
	if layer >= len(*result) {
		*result = append(*result, root.Val)
	}
	dfs(root.Right, layer+1, result)
	dfs(root.Left, layer+1, result)
	return
}
