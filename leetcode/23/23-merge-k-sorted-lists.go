/*
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/
package main

import (
	"fmt"
)

func createList(array []int) *ListNode {
	header := &ListNode{}
	tail := header
	for _, num := range array {
		tail.Next = &ListNode{Val: num}
		tail = tail.Next
	}
	return header.Next
}

func main() {
	// [[-6,-3,-1,1,2,2,2],[-10,-8,-6,-2,4],[-2],[-8,-4,-3,-3,-2,-1,1,2,3],[-8,-6,-5,-4,-2,-2,2,4]]
	// expect: [-10,-8,-8,-8,-6,-6,-6,-5,-4,-4,-3,-3,-3,-2,-2,-2,-2,-2,-1,-1,1,1,2,2,2,2,2,3,4,4]
	// output: [-10,-8,-8,-8,-6,-6,-5,-4,-6,-4,-3,-3,-3,-2,-2,-2,-2,-2,-1,-1,1,1,2,2,2,2,2,3,4,4]
	list1 := createList([]int{-6, -3, -1, 1, 2, 2, 2})
	list2 := createList([]int{-10, -8, -6, -2, 4})
	list3 := createList([]int{-2})
	list4 := createList([]int{-8, -4, -3, -3, -2, -1, 1, 2, 3})
	list5 := createList([]int{-8, -6, -5, -4, -2, -2, 2, 4})
	lists := mergeKLists([]*ListNode{list1, list2, list3, list4, list5})
	for h := lists; h != nil; h = h.Next {
		fmt.Println(h.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var validList []*ListNode
	for i := range lists {
		if lists[i] != nil {
			validList = append(validList, lists[i])
		}
	}
	lists = validList

	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	var header = &ListNode{}
	var tail = header
	minStack := minStack{
		array: make([]*ListNode, len(lists)+1),
		size:  1,
	}

	for i := range lists {
		minStack.Push(lists[i])
	}
	for {
		pop, _ := minStack.Pop()
		tail.Next = &ListNode{Val: pop.Val}
		tail = tail.Next
		if pop.Next != nil {
			minStack.Push(pop.Next)
		}
		if minStack.size == 1 {
			tail.Next = minStack.array[0]
			break
		}
	}
	return header.Next
}

type minStack struct {
	array []*ListNode
	size  int
}

func (p *minStack) Push(value *ListNode) bool {
	if p.size == len(p.array) {
		return false
	}
	arr := p.array
	pos := p.size
	p.size++
	for arr[pos] = value; pos > 1 && arr[pos].Val < arr[pos/2].Val; {
		arr[pos], arr[pos/2] = arr[pos/2], arr[pos]
		pos /= 2
	}
	return true
}

func (p *minStack) Pop() (*ListNode, bool) {
	if p.size <= 1 {
		return nil, false
	}
	arr := p.array
	arr[1], arr[p.size-1] = arr[p.size-1], arr[1]
	p.size--
	pos := 1
	for pos*2 < p.size {
		var minPos int
		if pos*2+1 >= p.size {
			minPos = pos * 2
		} else if arr[pos*2].Val < arr[pos*2+1].Val {
			minPos = pos * 2
		} else {
			minPos = pos*2 + 1
		}
		if arr[pos].Val > arr[minPos].Val {
			arr[pos], arr[minPos] = arr[minPos], arr[pos]
			pos = minPos
		} else {
			break
		}
	}
	return arr[p.size], true
}
