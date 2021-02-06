package leetcode350

// method of hash
func IntersectReview(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)

	numCount := make(map[int]int)
	for _, num := range nums1 {
		numCount[num]++
	}

	for _, num := range nums2 {
		if numCount[num] > 0 {
			numCount[num]--
			result = append(result, num)
		}
	}

	return result
}
