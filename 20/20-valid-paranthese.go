/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
*/
package main

import "container/list"

func main() {
	println(isValid("[()[]{[()]}]"))
}

func isValid(s string) bool {
	stack := list.New()
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack.PushBack(byte(c))
		case ')':
			if stack.Back() == nil || stack.Back().Value.(byte) != '(' {
				return false
			} else {
				stack.Remove(stack.Back())
			}
		case ']':
			if stack.Back() == nil || stack.Back().Value.(byte) != '[' {
				return false
			} else {
				stack.Remove(stack.Back())
			}
		case '}':
			if stack.Back() == nil || stack.Back().Value.(byte) != '{' {
				return false
			} else {
				stack.Remove(stack.Back())
			}
		}
	}
	if stack.Len() == 0 {
		return true
	}
	return false
}
