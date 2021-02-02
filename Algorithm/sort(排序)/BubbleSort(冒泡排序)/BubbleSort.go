package BubbleSort

func bubblesort(nums []int) {
	//get min
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

	// 冒泡排序
	sort := false
	for i := len(nums) - 1; i > 0 && sort == false; i-- {
		sort = true
		for j := 0; j < i; j++ {
			if IsMin(nums[j+1], nums[j]) { // j+1比j小，则交换位置
				sort = false
				swap(j+1, j)
			}
		}
	}
}
