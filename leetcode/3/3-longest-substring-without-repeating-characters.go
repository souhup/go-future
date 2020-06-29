/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

package main

func main() {
	var str = "pwwkew"
	println(lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(s string) int {
	var maxLength = 0
	var start int
	var charSet = make(map[byte]int)
	for pos, char := range s {
		if oldPos, ok := charSet[byte(char)]; ok && charSet[byte(char)] >= start {
			maxLength = max(maxLength, pos-start)
			start = oldPos + 1
		}
		charSet[byte(char)] = pos
	}
	maxLength = max(maxLength, len(s)-start)
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
