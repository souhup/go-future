/*
给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
说明:

可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n2) 。
进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
*/

package main

import "fmt"

func main() {
	data := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	fmt.Print(lengthOfLIS(data))
}

/**
解法1，动态规划
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var dp = make([]int, len(nums))
	for i := range nums {
		baseLen := 0
		for j := i; j >= 0; j-- {
			if nums[j] < nums[i] {
				if dp[j] > baseLen {
					baseLen = dp[j]
				}
			}
		}
		baseLen++
		dp[i] = baseLen
	}

	var result = 0
	for _, value := range dp {
		if value > result {
			result = value
		}
	}
	return result
}
*/

// 解法2，动态规划加二分法
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var dp = make([]int, 0)
	for _, value := range nums {
		tmpRes := searchTarget(dp, value)
		if tmpRes == -1 {
			dp = append(dp, value)
		} else {
			dp[tmpRes] = value
		}
	}
	return len(dp)
}

func searchTarget(nums []int, target int) int {
	if len(nums) == 0 || target > nums[len(nums)-1] {
		return -1
	}
	if target < nums[0] {
		return 0
	}

	left := 0
	right := len(nums)
	for left < right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			return middle
		} else if left == middle {
			return right
		} else if nums[middle] < target {
			left = middle
		} else {
			right = middle
		}
	}
	return -1
}
