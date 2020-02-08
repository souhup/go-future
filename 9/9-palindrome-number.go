/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。

*/
package main

func main() {
	println(isPalindrome(14641))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	numArr := make([]int, 0)
	for num := x; num > 0; num /= 10 {
		numArr = append(numArr, num%10)
	}
	for i := 0; i <= len(numArr)/2; i++ {
		if numArr[i] != numArr[len(numArr)-i-1] {
			return false
		}
	}
	return true
}
