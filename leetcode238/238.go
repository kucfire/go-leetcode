/*
	leetcode tag:238 title:除自身以外数组的乘积
*/

package leetcode238

// method of myself
func ProductExceptSelf(nums []int) []int {
	// 左侧元素乘积
	l := make([]int, len(nums))
	l[0] = 1
	for i := 1; i < len(nums); i++ {
		l[i] = l[i-1] * nums[i-1]
	}

	r := 1
	for i := len(nums) - 1; i >= 0; i-- {
		l[i] = l[i] * r
		r *= nums[i]
	}

	return l
}
