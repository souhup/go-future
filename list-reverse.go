// 问题：如何实现一个高效的单向链表逆序输出
package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func main() {
	var header *Node
	var last *Node

	header = &Node{data: 0}
	last = header
	for i := 1; i <= 10; i++ {
		nodeTmp := &Node{data: i}
		last.next = nodeTmp
		last = nodeTmp
	}
	reverse(header)
}

// 逆序输出
func reverse(node *Node) {
	if nil == node {
		return
	}
	fmt.Println(node.data)
	reverse(node.next)
}
