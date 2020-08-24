/*
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
说明：

所有输入均为小写字母。
不考虑答案输出的顺序。
*/
package main

import (
	"fmt"
)

func main() {
	data := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(data))
}

func groupAnagrams(strs []string) [][]string {
	var strMap = make(map[[26]int][]string)
	for _, str := range strs {
		sortStr := getStrKey(str)
		strMap[sortStr] = append(strMap[sortStr], str)
	}
	var result [][]string
	for _, value := range strMap {
		result = append(result, value)
	}
	return result
}

func getStrKey(str string) [26]int {
	var nums [26]int
	for _, s := range str {
		nums[s-'a']++
	}
	return nums
}
