/*
	leetcode tag:011 title:盛最多水的容器
*/

package leetcode011

func MaxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	maxCapacity, tmpCapacity := 0, 0
	left, right := 0, len(height)-1
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	for right > left {
		tmpCapacity = (right - left) * min(height[left], height[right])
		if tmpCapacity > maxCapacity {
			maxCapacity = tmpCapacity
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxCapacity
}
