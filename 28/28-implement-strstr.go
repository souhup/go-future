/*
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
*/

package main

func main() {
	src := "ababababca"
	dst := "abababca"
	println(strStr(src, dst))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	// 获取next数组
	next := getPartitionMathTable(needle)
	for i, j := 0, 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - j + 1
		}
	}
	return -1
}

func getPartitionMathTable(str string) []int {
	next := make([]int, len(str))
	for i, maxLength := 1, 0; i < len(str); i++ {
		for maxLength > 0 && str[i] != str[maxLength] {
			maxLength = next[maxLength-1]
		}
		if str[i] == str[maxLength] {
			maxLength++
		}
		next[i] = maxLength
	}
	return next
}
