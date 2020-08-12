/*
给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	//data := [][]int{{8, 10}, {1, 3}, {15, 18}, {2, 6}}
	//data := [][]int{{8, 10}, {10, 19}}
	data := [][]int{{1, 4}, {2, 3}}
	fmt.Println(merge(data))
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		if len(result) == 0 || len(result) > 0 && result[len(result)-1][1] < intervals[i][0] {
			result = append(result, intervals[i])
		} else {
			result[len(result)-1][1] = max(result[len(result)-1][1], intervals[i][1])
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
