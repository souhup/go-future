/*
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/
package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{Val: 4},
		}}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{Val: 4,
				Next: &ListNode{Val: 5},
			},
		}}

	l3 := mergeTwoLists(l1, l2)

	for head := l3; head != nil; head = head.Next {
		fmt.Println(head.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var head = &ListNode{}
	var last = head
	var lastlast *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			last.Val = l1.Val
			l1 = l1.Next
		} else {
			last.Val = l2.Val
			l2 = l2.Next
		}
		lastlast = last
		last.Next = &ListNode{}
		last = last.Next
	}
	if l1 != nil && lastlast != nil {
		lastlast.Next = l1
	} else if l2 != nil && lastlast != nil {
		lastlast.Next = l2
	} else if lastlast != nil {
		lastlast.Next = nil
	}
	return head
}
