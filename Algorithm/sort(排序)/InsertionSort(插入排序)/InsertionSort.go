package InsertionSort

func insertionsort(nums []int) {
	// 判断更小的值
	IsMin := func(x, y int) bool {
		if x >= y {
			return false
		}
		return true
	}

	// 交换位置
	swap := func(x, y int) {
		nums[x], nums[y] = nums[y], nums[x]
	}

	for i := 1; i < len(nums); i++ {
		for j := i; j > 0 && IsMin(nums[j], nums[j-1]); j-- {
			swap(nums[j], nums[j-1])
		}
	}
}
