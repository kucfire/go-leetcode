/*
	leetcode tag:014 title:最长公共前缀
*/

package leetcode014

import "math"

// LongestCommonPrefix : normal method
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var result string
	min := func(strs []string) int {
		minLength := math.MaxInt32
		for _, v := range strs {
			if len(v) < minLength {
				minLength = len(v)
			}
		}
		return minLength
	}

	// 获取数组内的元素中，最短的长度
	minLength := min(strs)
loop1:
	for i := 0; i < minLength; i++ {
		var tmp byte
		for j, v := range strs {
			if j == 0 {
				tmp = v[i]
				continue
			}
			if v[i] != tmp {
				break loop1
			}
		}
		result += string(tmp)
	}
	return result
}
