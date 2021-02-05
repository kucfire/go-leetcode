package leetcode217

func ContainsDuplicateReview(nums []int) bool {
	numArray := make(map[int]int)
	for _, num := range nums {
		numArray[num]++
		if numArray[num] > 1 {
			return true
		}
	}
	return false
}
