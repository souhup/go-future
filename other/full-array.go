// 题目: 打印生成1到N的全排列
package main

import "fmt"

func main() {
	fullArray(3)
}

func fullArray(n int) {
	array := make([]int, n)
	for i := range array {
		array[i] = i
	}
	dfs(array, n, 0)
}

func dfs(array []int, max, position int) {
	if max == position {
		fmt.Println(array)
		return
	}
	for i := position; i < max; i++ {
		array[i], array[position] = array[position], array[i]
		dfs(array, max, position+1)
		array[i], array[position] = array[position], array[i]
	}
}
