/*
	leetcode tag:198 title:打家劫舍
*/

package leetcode198

// Rob :
func Rob(nums []int) int {
	// 初始条件判断，len(nums)必须大于1
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	// var result int

	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}

	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}

	return second
}
