/*
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。

请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。


示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5
*/
package main

import (
	"fmt"
)

func main() {
	nums1 := []int{4}
	nums2 := []int{1, 2, 3, 5, 6}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen%2 == 1 {
		return findKInTwoArr(1+(totalLen-1)/2, nums1, nums2)
	} else {
		return (findKInTwoArr(totalLen/2, nums1, nums2) + findKInTwoArr(totalLen/2+1, nums1, nums2)) / 2
	}
}

// 寻找两个有序数组中最小的第K个数字
func findKInTwoArr(k int, nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return float64(nums2[k-1])
	}
	if len(nums2) == 0 {
		return float64(nums1[k-1])
	}
	if k == 1 {
		if nums1[0] < nums2[0] {
			return float64(nums1[0])
		} else {
			return float64(nums2[0])
		}
	}
	if len(nums1) < k/2 {
		if nums1[len(nums1)-1] < nums2[k/2-1] {
			return findKInTwoArr(k-len(nums1), nil, nums2)
		} else {
			return findKInTwoArr(k-k/2, nums1, nums2[k/2:])
		}
	}
	if len(nums2) < k/2 {
		if nums2[len(nums2)-1] < nums1[k/2-1] {
			return findKInTwoArr(k-len(nums2), nums1, nil)
		} else {
			return findKInTwoArr(k-k/2, nums1[k/2:], nums2)
		}
	}
	if nums1[k/2-1] < nums2[k/2-1] {
		return findKInTwoArr(k-k/2, nums1[k/2:], nums2)
	} else {
		return findKInTwoArr(k-k/2, nums1, nums2[k/2:])
	}
}
