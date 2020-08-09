/*

给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
*/
package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	// dp[i]表示以i结尾的连续子数组和
	// dp[i] = max(dp[i-1] + nums[i], nums[i])
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxSubArray2(nums []int) int {
	// dp[i]表示以i结尾的连续子数组和
	// dp[i] = max(dp[i-1] + nums[i], nums[i])
	if len(nums) == 0 {
		return 0
	}
	dp := nums[0]
	result := dp
	for i := 1; i < len(nums); i++ {
		dp = max(dp+nums[i], nums[i])
		if dp > result {
			result = dp
		}
	}
	return result
}
