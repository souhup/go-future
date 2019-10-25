// 题目：给定一个二叉搜索树(BST)，找到树中第 K 小的节点。
package main

import (
	"fmt"
)

// BST is binary search tree
type BST struct {
	key, value  int
	left, right *BST
}

func main() {
	tree := createBST()
	fmt.Println(find(tree, 3))
}

// 查找bst第k个节点的值，未找到就返回0
func find(b *BST, k int) (value int) {
	if k < 0 {
		return
	}
	_, value = findRecursive(b, &k)
	return
}

// countRecurisive 对bst使用中序遍历
// 用计数方式控制退出遍历，参数c就是已遍历节点数
func findRecursive(b *BST, k *int) (ok bool, value int) {
	if b == nil {
		return
	}
	ok, value = findRecursive(b.left, k)
	if ok {
		return
	}

	*k--
	if *k == 0 {
		return true, b.value
	}

	ok, value = findRecursive(b.right, k)
	if ok {
		return
	}
	return
}

func (bst *BST) setLeft(b *BST) {
	bst.left = b
}

func (bst *BST) setRight(b *BST) {
	bst.right = b
}

func createBST() *BST {
	b1 := &BST{key: 1, value: 10}
	b2 := &BST{key: 2, value: 20}
	b3 := &BST{key: 3, value: 30}
	b4 := &BST{key: 4, value: 40}
	b5 := &BST{key: 5, value: 50}
	b6 := &BST{key: 6, value: 60}
	b7 := &BST{key: 7, value: 70}
	b8 := &BST{key: 8, value: 80}
	b9 := &BST{key: 9, value: 90}

	b5.setLeft(b3)
	b5.setRight(b7)
	b3.setLeft(b2)
	b3.setRight(b4)
	b2.setLeft(b1)
	b7.setLeft(b6)
	b7.setRight(b8)
	b8.setRight(b9)

	return b5
}
