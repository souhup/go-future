/*
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:

输入: 4
输出: 2
示例 2:

输入: 8
输出: 2
说明: 8 的平方根是 2.82842...,
     由于返回类型是整数，小数部分将被舍去。
*/

package main

func main() {
	println(mySqrt(100000))
}

func mySqrt(x int) int {
	switch x {
	case 0:
		return 0
	case 1:
		return 1
	}
	left := 1
	right := x
	for {
		mid := left + (right-left)/2
		multi := mid * mid
		if multi == x || left == mid {
			if right*right < x {
				return right
			} else {
				return mid
			}
		} else if multi > x {
			right = mid - 1
		} else {
			left = mid
		}
	}
}
