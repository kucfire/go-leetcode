package leetcode169

import "sort"

func MajorityElementReview(nums []int) int {
	// 先排序
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func MajorityElementReviewStack(nums []int) int {
	stack := make([]int, 0, len(nums))
	for _, num := range nums {
		if len(stack) > 0 && stack[len(stack)-1] != num {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, num)
	}
	return stack[0]
}
