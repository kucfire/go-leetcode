package leetcode283

func MoveZeroesReview(nums []int) {
	for move, i := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[move] = nums[move], nums[i]
			move++
		}
	}
}
