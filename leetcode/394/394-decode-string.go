/*
给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
*/
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(decodeString2("1[b2[c]]"))
}

func decodeString(s string) string {
	data := []byte(s)
	result, _ := decode(data, 0)
	return string(result)
}

func decode(data []byte, i int) ([]byte, int) {
	buffer := bytes.NewBuffer(nil)
	buffer.Grow(len(data))
	k := 0
	for ; i < len(data); i++ {
		if (data[i] >= 'a' && data[i] <= 'z') || (data[i] >= 'A' && data[i] <= 'Z') {
			buffer.WriteByte(data[i])
		} else if data[i] >= '0' && data[i] <= '9' {
			k = k*10 + int(data[i]-'0')
		} else if data[i] == '[' {
			result, nextI := decode(data, i+1)
			buffer.Write(bytes.Repeat(result, k))
			k = 0
			i = nextI
		} else {
			return buffer.Bytes(), i
		}
	}
	return buffer.Bytes(), len(data)
}

func decodeString2(s string) string {
	var strStack = make([]string, 0)
	var numStack = make([]int, 0)

	var tmpStr string
	var k int
	for _, b := range []byte(s) {
		if b >= '0' && b <= '9' {
			k = k*10 + int(b-'0')
		} else if b == '[' {
			numStack = append(numStack, k)
			k = 0
			strStack = append(strStack, tmpStr)
			tmpStr = ""
		} else if b == ']' {
			count := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			preStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			tmpStr = preStr + strings.Repeat(tmpStr, count)
		} else {
			tmpStr += string(b)
		}
	}
	return tmpStr
}
