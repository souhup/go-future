package main

/**
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

func main() {
	var l1 = &ListNode{
		Val: 9,
		Next: &ListNode{
			Val:  9,
			Next: &ListNode{Val: 9},
		},
	}
	var l2 = &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 3,
		},
	}
	result := addTwoNumbers(l1, l2)
	for cur := result; cur != nil; cur = cur.Next {
		println(cur.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result = l1

	for l1.Next != nil && l2.Next != nil {
		l1.Val += l2.Val
		l1 = l1.Next
		l2 = l2.Next
	}
	l1.Val += l2.Val

	if l2.Next != nil {
		l1.Next = l2.Next
	}

	for cur := result; cur != nil; cur = cur.Next {
		if cur.Val >= 10 {
			cur.Val %= 10
			if cur.Next != nil {
				cur.Next.Val++
			} else {
				cur.Next = &ListNode{Val: 1}
			}
		}
	}
	return result
}
