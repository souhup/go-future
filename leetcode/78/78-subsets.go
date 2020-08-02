/*
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

package main

import "fmt"

func main() {
	data := []int{1, 2, 3}
	fmt.Println(subsets(data))
}

var result [][]int

func subsets(nums []int) [][]int {
	result = make([][]int, 0)
	dfs(nums, []int{})
	return result
}

func dfs(nums []int, queue []int, ) {
	if len(nums) == 0 {
		result = append(result, append([]int{}, queue...))
		return
	}
	// 弃用
	dfs(nums[1:], queue)
	// 选用
	queue = append(queue, nums[0])
	dfs(nums[1:], queue)
	queue = queue[:len(queue)-1]
	return
}
