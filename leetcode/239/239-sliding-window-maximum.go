/*
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。



进阶：

你能在线性时间复杂度内解决此题吗？



示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7


提示：

1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
1 <= k <= nums.length
*/
package main

import (
	"fmt"
)

func main() {
	//nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	nums := []int{1, 3}
	fmt.Println(maxSlidingWindow(nums, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k < 1 {
		return nil
	}
	if len(nums) < k {
		max := 0
		for _, num := range nums {
			if num > max {
				max = num
			}
		}
		return []int{max}
	}

	result := make([]int, len(nums)-k+1)
	array := make([]int, 0, len(nums))
	for i, num := range nums {
		// 尝试删除头部
		if len(array) > 0 && array[0] < i-k+1 {
			array = array[1:]
		}

		// 尝试追加尾部
		for len(array) > 0 && nums[array[len(array)-1]] < num {
			array = array[:len(array)-1]
		}
		array = append(array, i)
		if i+1 >= k {
			result[i-k+1] = nums[array[0]]
		}
	}
	return result
}
