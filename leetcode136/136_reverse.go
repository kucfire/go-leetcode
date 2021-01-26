package leetcode136

// bit
func SingleNumberBit(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}
	return nums[0]
}

// hashmap
func SingleNumberHashMap(nums []int) int {
	arrayMap := make(map[int]int)
	for _, num := range nums {
		arrayMap[num]++
	}
	// var bigger int
	// var result int
	for key, value := range arrayMap {
		if value == 1 {
			return key
		}
	}
	return 0
}
