/*
给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

示例  1:

输入: "Let's take LeetCode contest"
输出: "s'teL ekat edoCteeL tsetnoc"
注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest") == "s'teL ekat edoCteeL tsetnoc")
}

func reverseWords(s string) string {
	bArr := bytes.Split([]byte(s), []byte(" "))
	for it := range bArr {
		reverse(bArr[it])
	}
	return string(bytes.Join(bArr, []byte(" ")))
}

func reverse(s []byte) []byte {
	for start, end := 0, len(s)-1; start < end; start, end = start+1, end-1 {
		s[start], s[end] = s[end], s[start]
	}
	return s
}
