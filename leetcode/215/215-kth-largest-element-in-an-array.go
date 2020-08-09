/*
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
说明:

你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
*/
package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(nums, 1))
}

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	stack := minStack{size: 1, array: make([]int, k+1)}
	stack.array[1] = -1 << 31
	for _, num := range nums {
		stack.push(num)
	}
	return stack.array[1]
}

type minStack struct {
	size  int
	array []int
}

func (m *minStack) push(value int) {
	if m.size == len(m.array) && value <= m.array[1] {
		return
	}
	if m.size == len(m.array) {
		m.pop()
	}
	arr := m.array
	arr[m.size] = value
	pos := m.size
	m.size++
	for pos > 1 {
		if arr[pos] >= arr[pos/2] {
			break
		}
		arr[pos], arr[pos/2] = arr[pos/2], arr[pos]
		pos /= 2
	}
}

func (m *minStack) pop() int {
	arr := m.array
	arr[1], arr[m.size-1] = arr[m.size-1], arr[1]
	m.size--
	pos := 1
	for 2*pos < m.size {
		var minPos int
		if 2*pos+1 < m.size && arr[2*pos+1] < arr[2*pos] {
			minPos = 2*pos + 1
		} else {
			minPos = 2 * pos
		}
		if arr[pos] <= arr[minPos] {
			break
		}
		arr[pos], arr[minPos] = arr[minPos], arr[pos]
		pos = minPos
	}
	return arr[m.size]
}
