/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(-1563847412))
}

func reverse(x int) int {
	flag := 1
	if x < 0 {
		flag = -1
		x = x * -1
	}
	result := 0
	for num := x; num > 0; num /= 10 {
		result = result*10 + num%10
	}
	if flag*result > math.MaxInt32 || flag*result < math.MinInt32 {
		return 0
	} else {
		return flag * result
	}
}
