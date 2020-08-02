/*
给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

注意:

每个数组中的元素不会超过 100
数组的大小不会超过 200
示例 1:

输入: [1, 5, 11, 5]

输出: true

解释: 数组可以分割成 [1, 5, 5] 和 [11].

示例2:

输入: [1, 2, 3, 5]

输出: false
*/
package main

import "fmt"

func main() {
	data := []int{1, 5, 11, 5}
	fmt.Println(canPartition2(data))
}

func canPartition(nums []int) bool {
	total := 0
	for _, num := range nums {
		total += num
	}
	// 如果是奇数, 则直接返回
	if total&1 == 1 {
		return false
	}

	// dp[i][j]表示前i个数字的集合是能够满足任意数字之和为j. 满足则为true, 否则为false
	var dp = make([][]bool, len(nums)+1)
	for i := range dp {
		dp[i] = make([]bool, total/2+1)
	}

	// 对于背包为0的，全部为true
	for i := range dp {
		dp[i][0] = true
	}

	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= total/2; j++ {
			if j < nums[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				// 不选择数字的话,状态为dp[i][j] = dp[i-1][j]
				// 选择数字的话, 状态为dp[i][j] = dp[i-1][j-nums[i-1]]
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
		// 剪支, 如果至少第二维为total/2且为true,接可以直接返回结果
		if dp[i][total/2] {
			return true
		}
	}
	for i := range dp {
		fmt.Println(dp[i])
	}
	return dp[len(nums)][total/2]
}

// 数组压缩
func canPartition2(nums []int) bool {
	total := 0
	for _, num := range nums {
		total += num
	}
	// 如果是奇数, 则直接返回
	if total&1 == 1 {
		return false
	}

	// dp[i]表示前i个数字的集合是能够满足任意数字之和为j. 满足则为true, 否则为false, 此为二维数组的压缩
	var dp = make([]bool, total/2+1)
	dp[0] = true

	for i := 1; i <= len(nums); i++ {
		for j := total / 2; j >= 1; j-- {
			if j >= nums[i-1] {
				// 不选择数字的话,状态为dp[i][j] = dp[i-1][j]
				// 选择数字的话, 状态为dp[i][j] = dp[i-1][j-nums[i-1]]
				dp[j] = dp[j] || dp[j-nums[i-1]]
			}
		}
		// 剪支, 如果至少第二维为total/2且为true,接可以直接返回结果
		if dp[total/2] {
			return true
		}
	}
	return dp[total/2]
}
