/*

给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？
*/
package main

import "fmt"

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	newHead := removeNthFromEnd(head, 4)
	for i := newHead; i != nil; i = i.Next {
		fmt.Println(i.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	k := remove(head, n)
	if k == n {
		head = head.Next
	}
	return head
}

func remove(head *ListNode, n int) int {
	if head.Next == nil {
		return 1
	}
	k := remove(head.Next, n)
	if k == n {
		head.Next = head.Next.Next
	}
	return k + 1
}
