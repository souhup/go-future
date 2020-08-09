/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 0, 0}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result = make([][]int, 0)
	for i, num := range nums {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		if num > 0 {
			break
		}
		left, right := i+1, len(nums)-1
		for left < right {
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if right < len(nums)-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			if nums[left]+nums[right] == -num {
				result = append(result, []int{num, nums[left], nums[right]})
				left, right = left+1, right-1
				continue
			} else if nums[left]+nums[right] < -num {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
