/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6
*/
package main

import "fmt"

func main() {
	height := []int{2, 1, 2}
	fmt.Println(trap(height))
}

func trap(height []int) int {
	fullHeight := make([]int, len(height))
	for maxHeight, i := 0, 0; i < len(height); i++ {
		maxHeight = max(maxHeight, height[i])
		fullHeight[i] = maxHeight
	}
	for maxHeight, i := 0, len(height)-1; i >= 0; i-- {
		maxHeight = max(maxHeight, height[i])
		fullHeight[i] = min(maxHeight, fullHeight[i])
	}
	var result = 0
	for i := 0; i < len(height); i++ {
		result += fullHeight[i] - height[i]
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

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
