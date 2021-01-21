/*
	leetcode tag:055 title:跳跃游戏
*/

package leetcode055

// method of normal
func CanJump(nums []int) bool {
	target := len(nums)
	// 无法完成
	if target < 1 {
		return false
	}
	// 自己跳自己
	if target == 1 {
		return true
	}
	// 一开始就跳不动了
	if nums[0] == 0 {
		return false
	}

	// 代表
	zero := -1
	for i := target - 2; i >= 0; i-- {
		if zero > 0 {
			if i+nums[i] > zero {
				zero = -1
			}
			continue
		}

		if nums[i] == 0 {
			zero = i
		}
	}

	if zero < 0 {
		return true
	}
	return false
}
