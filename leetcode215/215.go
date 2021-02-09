/*
	leetcode tag:215 title:数组中第k大元素
*/

package leetcode215

import "sort"

func FindKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// 自己写排序(性能极差)
func FindKthLargestMyself(nums []int, k int) int {
	var Sorts func(nums []int)
	Sorts = func(nums []int) {
		if len(nums)-1 == 0 {
			return
		}

		for i := 1; i < len(nums); i++ {
			if nums[i] <= nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		}
		Sorts(nums[:len(nums)-1])
	}
	Sorts(nums)
	return nums[len(nums)-k]
}
