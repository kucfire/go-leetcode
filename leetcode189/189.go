/*
	leetcode tag:189 title:旋转数组
*/

package leetcode189

// Rotate : normal method
func Rotate(nums []int, k int) {
	for ; k > 0; k-- {
		tmp := nums[len(nums)-1]
		copy(nums[1:], nums[:len(nums)-1])
		nums[0] = tmp
	}
}

// Rotate2 : the most fast method
func Rotate2(nums []int, k int) {
	k = k % len(nums)
	copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
}
