/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。
*/
package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLen := 1<<31 - 1
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}
	resultLen := 0
	for i := 0; i < minLen; i++ {
		tmpChar := strs[0][i]
		for _, str := range strs {
			if str[i] != tmpChar {
				return strs[0][:resultLen]
			}
		}
		resultLen++
	}
	return strs[0][:resultLen]
}
