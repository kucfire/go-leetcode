package SelectionSort

func selectsort(nums []int) {
	// get min
	getMin := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	// change position
	swap := func(x, y int) {
		nums[x], nums[y] = nums[y], nums[x]
	}

	// select sort
	for i := 0; i < len(nums); i++ {
		min := nums[i]
		// 找出最小值
		for j := i + 1; j < len(nums); j++ {
			min = getMin(min, nums[j])
		}
		swap(i, min)
	}
}
