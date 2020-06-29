// 题目：已知 sqrt (2)约等于 1.414，要求不用数学库，求 sqrt (2)精确到小数点后 10 位。
package main

import (
	"fmt"
	"math"
)

const EPSILON = 0.000_000_000_1

func main() {
	binarySqrt(2)
	newtonSqrt(2)
}

// 二分法
func binarySqrt(target float64) {
	upperBound := target
	lowBound := 1.0
	value := (upperBound + lowBound) / 2
	for {
		accumulate := value * value
		if math.Abs(accumulate-target) <= EPSILON {
			break
		}
		if accumulate > target {
			upperBound = value
		} else {
			lowBound = value
		}
		value = (upperBound + lowBound) / 2
	}
	fmt.Println("二分法\t\t", value)
}

// 牛顿迭代法
// L的切线方程为 y=f(x0)+f′(x0)(x−x0)
// x1的一次近似值 x1=x0−f(x0)/f′(x0)
// 开根号的问题可以看作求解f(x)=x^2−a=0的根
//
func newtonSqrt(target float64) {
	value := 1.0
	for math.Abs(value*value-target) > EPSILON {
		value = (value + target/value) / 2
	}
	fmt.Println("牛顿切线法\t", value)
}
