// review
package leetcode088

import "sort"

func MergeReview(nums1 []int, m int, nums2 []int, n int) []int {
	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
	return nums1
}
