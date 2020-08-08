/*
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/
package main

import "fmt"

func main() {
	data := []int{1, 2, 3}
	result := permute(data)
	fmt.Println(result)
}

func permute(nums []int) [][]int {
	var result = make([][]int, 0)
	myPermute(&result, nums, 0)
	return result
}

func myPermute(result *[][]int, nums []int, level int) {
	if level == len(nums)-1 {
		*result = append(*result, append([]int{}, nums...))
	}
	for i := level; i < len(nums); i++ {
		nums[level], nums[i] = nums[i], nums[level]
		myPermute(result, nums, level+1)
		nums[level], nums[i] = nums[i], nums[level]
	}
}
