/*
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

示例：

给你这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5

说明：

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*/
package main

import (
	"fmt"
)

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	newHead := reverseKGroup(head, 3)
	for h := newHead; h != nil; h = h.Next {
		fmt.Println(h.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	var result = &ListNode{}
	var newHead = result

	var stack = make([]*ListNode, k)
	var stackPtr = 0
	for iter := head; iter != nil; {
		stack[stackPtr] = iter
		stackPtr++
		if stackPtr >= k {
			iter = iter.Next
			for ; stackPtr > 0; stackPtr-- {
				newHead.Next = stack[stackPtr-1]
				newHead = stack[stackPtr-1]
			}
			newHead.Next = nil
		} else if iter.Next == nil && stackPtr > 0 {
			newHead.Next = stack[0]
			iter = iter.Next
		} else {
			iter = iter.Next
		}
	}
	return result.Next
}
