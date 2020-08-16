/*
在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。

示例:

输入:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

输出: 4
*/
package main

import "fmt"

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '1', '1', '0'},
	}

	fmt.Println(maximalSquare(matrix))
}

func maximalSquare(matrix [][]byte) int {
	//dp[i][j] =  min(dp[i-1]dp[j-1], dp[i][j-1], dp[i-1][j]) + 1, if dp[i][j] == 1
	var dp = make([][]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
	}

	var result = 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 || j == 0 || matrix[i][j] == '0' {
				dp[i][j] = int(matrix[i][j] - '0')
				result = max(result, dp[i][j])
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
				result = max(result, dp[i][j])
			}
		}
	}
	return result * result
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}
