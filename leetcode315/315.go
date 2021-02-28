/*
	leetcode tag:315 title:计算右侧小于当前元素的个数
*/

package leetcode315

func CountSmaller(nums []int) []int {
	dp := make([]int, len(nums))
	if len(nums) == 0 {
		return dp
	}

	dp[len(nums)-1] = 0
	//从后往前更新
	for i := len(nums) - 2; i >= 0; i-- {
		count := 0
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				count++
			}
		}
		dp[i] = count
	}
	return dp
}
