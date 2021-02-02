package ShellSort

func shellsort(nums []int) { // 判断更小的值
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

	N := len(nums)
	h := 1

	for h < N/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j >= h && IsMin(nums[j], nums[j-h]); j -= h {
				swap(nums[i], nums[j])
			}
		}
		h = h / 3
	}
}
