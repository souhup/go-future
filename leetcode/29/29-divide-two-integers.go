/*
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2



示例 1:

输入: dividend = 10, divisor = 3
输出: 3
解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
解释: 7/-3 = truncate(-2.33333..) = -2


提示：

被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。
*/
package main

import "fmt"

func main() {
	fmt.Println(divide(6666666667, 2))
}

func divide(dividend int, divisor int) int {
	// 特殊情况
	if dividend == -1<<31 && divisor == -1 {
		return 1<<31 - 1
	}
	if divisor == 1 {
		return dividend
	} else if divisor == -1 {
		return -dividend
	}

	var positiveFlag bool
	if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 {
		positiveFlag = true
	}
	if dividend > 0 {
		dividend = -dividend
	}
	if divisor > 0 {
		divisor = - divisor
	}

	result := quotient(dividend, divisor)
	if positiveFlag {
		return result
	} else {
		return -result
	}
}

// dividend, divisor 均不大于0
func quotient(dividend, divisor int) int {
	if dividend > divisor {
		return 0
	} else if dividend <= divisor && dividend > divisor<<1 {
		return 1
	}
	i := 0
	for ; dividend < divisor<<(i+1); i++ {
	}
	return 1<<i + quotient(dividend-divisor<<i, divisor)
}
