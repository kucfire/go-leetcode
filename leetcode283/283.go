/*
	leetcode tag:283 title:移动零
*/

package leetcode283

// MoveZeroes : normal method
func MoveZeroes(nums []int) {
	zeroStart := 0
	tmp := 0
	for i, num := range nums {
		if num != 0 {
			tmp = num
			nums[i] = 0
			nums[zeroStart] = tmp
			zeroStart++
		}
	}
}
