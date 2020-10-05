/*

给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。

示例：

输入：
A: [1,2,3,2,1]
B: [3,2,1,4,7]
输出：3
解释：
长度最长的公共子数组是 [3, 2, 1] 。


提示：

1 <= len(A), len(B) <= 1000
0 <= A[i], B[i] < 100
*/
package main

import "fmt"

func main() {
	s1 := []int{0, 1, 1, 1}
	s2 := []int{1, 0, 1, 0, 1}
	fmt.Println(findLength(s1, s2))
}

func findLength(A []int, B []int) int {
	// dp[i][j] 表示 A中以i结尾的字符串, B中以j结尾的字符串
	// dp[i][j] = dp[i-1][j-1] + 1,  if A[i] == B[j]
	dp := make([][]int, len(A)+1)
	for i := range dp {
		dp[i] = make([]int, len(B)+1)
	}
	maxLen := 0
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				maxLen = max(maxLen, dp[i][j])
			}
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
