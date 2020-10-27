/*
	leetcode tag:350 title:两个数组的交集
*/

package leetcode350

import "sort"

// Intersect : method of sort
func Intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	result := make([]int, 0)
	start := 0
	for _, num := range nums1 {
	loop2:
		for i := start; i < len(nums2); i++ {
			if num == nums2[i] {
				result = append(result, num)
				start = i + 1
				break loop2
			}
		}
	}
	return result
}
