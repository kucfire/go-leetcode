/*
	leetcode tag:088 title:合并两个有序数组
*/

package leetcode088

import "sort"

// normal method
func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	for i, j := m, 0; i < m+n && j < n; i, j = i+1, j+1 {
		nums1[i] = nums2[j]
	}
	sort.Ints(nums1)
	return nums1
}
