package leetcode189

func RotateReview(nums []int, k int) {
	// 分步执行
	for ; k > 0; k-- {
		tmp := nums[len(nums)-1]
		copy(nums[1:], nums[:len(nums)-1])
		nums[0] = tmp
	}
}
