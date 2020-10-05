/*
给定一个字符串 s ，找到其中最长的回文子序列，并返回该序列的长度。可以假设 s 的最大长度为 1000 。



示例 1:
输入:

"bbbab"
输出:

4
一个可能的最长回文子序列为 "bbbb"。

示例 2:
输入:

"cbbd"
输出:

2
一个可能的最长回文子序列为 "bb"。



提示：

1 <= s.length <= 1000
s 只包含小写英文字母
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
}

func longestPalindromeSubseq(s string) int {
	byteArr := []byte(s)
	if len(byteArr) == 0 {
		return 0
	}
	db := make([][]int, len(byteArr))
	for i := range db {
		db[i] = make([]int, len(byteArr))
	}

	// db[i][j] = db[i+1][j-1] + 2, if db[i] == db[j]
	// db[i][j] = max(db[i][j-1], db[i+1][j])
	for pos := 0; pos < len(byteArr); pos++ {
		for i, j := 0, pos; i < len(byteArr) && j < len(byteArr); i, j = i+1, j+1 {
			if i == j {
				db[i][j] = 1
			} else if byteArr[i] == byteArr[j] {
				db[i][j] = db[i+1][j-1] + 2
			} else {
				db[i][j] = max(db[i+1][j], db[i][j-1])
			}
		}
	}
	return db[0][len(byteArr)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
