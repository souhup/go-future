/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
*/
package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("bb"))
}

// 中心扩散法
func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	var maxPaliStart, maxPaliLen int
	for i := 0; i < len(s)-1; i++ {
		paliLen1 := checkPaliLen(s, i, i)
		paliLen2 := checkPaliLen(s, i, i+1)
		if paliLen1 > maxPaliLen {
			maxPaliLen = paliLen1
			maxPaliStart = i - (maxPaliLen-1)/2
		}
		if paliLen2 > maxPaliLen {
			maxPaliLen = paliLen2
			maxPaliStart = i - (maxPaliLen-2)/2
		}
	}
	return s[maxPaliStart : maxPaliStart+maxPaliLen]
}

func checkPaliLen(str string, start, end int) (paliLen int) {
	for start >= 0 && end < len(str) && str[start] == str[end] {
		start--
		end++
	}
	return end - start + 1 - 2
}
