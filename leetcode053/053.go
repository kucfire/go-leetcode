/*
	leetcode tag:053 title:最大子序和
*/

package leetcode053

// MaxSubArray : dynamic programming
func MaxSubArray(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		// 如果前面的自序和跟当前的子元素相加大于当前的子元素，将和存在当前对应下表的数组中
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > result {
			result = nums[i]
		}
	}
	return result
}
