package leetcode347

func TopKFrequent_Reverse(nums []int, k int) []int {
	arrayMap := make(map[int]int)
	var arrayList []int
	for _, num := range nums {
		_, ok := arrayMap[num]
		if !ok {
			arrayMap[num] = 1
			arrayList = append(arrayList, num)
		} else {
			arrayMap[num]++
		}
	}

	// 快速排序
	var quickSort func(arrayList []int, arrayMap map[int]int, left, right int)
	quickSort = func(arrayList []int, arrayMap map[int]int, left, right int) {
		if left >= right {
			return
		}

		// 取最后一个值作为中间值
		middle := arrayMap[arrayList[right]]
		// 从left左侧开始计数， left-1
		start := left - 1
		// 退出条件是到达middle所在位置的前一位，不能跟middle所在位置进行比较
		for i := left; i < right; i++ {
			if arrayMap[arrayList[i]] > middle { // 由于题干是要取前k个高频元素，所以数组需要从大到小排列
				start++
				arrayList[i], arrayList[start] = arrayList[start], arrayList[i]
			}
		}

		// 便利完成后，现将右侧的数跟当前数据的中位数交换一下
		arrayList[start+1], arrayList[right] = arrayList[right], arrayList[start+1]

		quickSort(arrayList, arrayMap, left, start)
		quickSort(arrayList, arrayMap, start+2, right) // +2是为了跳过中间值
	}

	quickSort(arrayList, arrayMap, 0, len(arrayList)-1)
	return arrayList[:k]
}

func TopKFrequent_Reverse2(nums []int, k int) []int {
	arrayMap := make(map[int]int)
	arrayList := make([]int, 0)
	for _, num := range nums {
		_, ok := arrayMap[num]
		if !ok {
			arrayList = append(arrayList, num)
		}
		arrayMap[num]++
	}

	var fastSort func(arrayList []int, arrayMap map[int]int, left, right int)
	fastSort = func(arrayList []int, arrayMap map[int]int, left, right int) {
		if left >= right {
			return
		}

		middle := arrayMap[arrayList[right]]
		start := left - 1
		for i := left; i < right; i++ {
			if arrayMap[arrayList[i]] > middle {
				start++
				arrayList[i], arrayList[start] = arrayList[start], arrayList[i]
			}
		}

		arrayList[start+1], arrayList[right] = arrayList[right], arrayList[start+1]

		fastSort(arrayList, arrayMap, left, start)
		fastSort(arrayList, arrayMap, start+2, right)
	}

	fastSort(arrayList, arrayMap, 0, len(arrayList)-1)

	return arrayList[:k]
}
