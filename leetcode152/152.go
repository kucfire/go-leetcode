/*
	leetcode tag:152 title:乘积最大子数组
*/

package leetcode152

import "math"

// MaxProductDP : method of DP
func MaxProductDP(nums []int) int {
	// result用作返回,并承接最大值的比较
	result := math.MinInt32
	n := len(nums)
	getMax := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	var dp func(start int) int
	dp = func(start int) int {
		if start >= n {
			return math.MinInt32
		}
		max := nums[start]
		total := nums[start]
		for i := start + 1; i < n; i++ {
			total = total * nums[i]
			if max < total {
				max = total
			}
		}
		max2 := dp(start + 1)
		return getMax(max, max2)
	}

	result = dp(0)
	return result
}

//
func MaxProduct(nums []int) int {
	max, min, result := nums[0], nums[0], nums[0]

	getMax := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	getMin := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	for i := 1; i < len(nums); i++ {
		mx, mn := max, min
		max = getMax(mx*nums[i], getMax(nums[i], mn*nums[i]))
		min = getMin(mn*nums[i], getMin(nums[i], mx*nums[i]))
		result = getMax(max, result)
	}
	return result
}
