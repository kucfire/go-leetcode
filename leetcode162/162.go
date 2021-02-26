/*
	leetcode tag:162 title:寻找峰值
*/

package leetcode162

// FindPeakElement : method of myself
func FindPeakElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}
	if len(nums) == 2 {
		max := func(i, j int) int {
			if nums[i] > nums[j] {
				return i
			}
			return j
		}
		return max(0, 1)
	}
	max := 0
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[max] {
			max = i
		}
		if nums[i-1] < nums[i] && nums[i+1] < nums[i] {
			return i
		}
	}

	if nums[len(nums)-1] > nums[max] {
		max = len(nums) - 1
	}
	return max
}

func FindPeakElement2(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums) - 1
}
