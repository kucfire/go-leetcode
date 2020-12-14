package leetcode347

func TopKFrequent(nums []int, k int) []int {
	// 计数
	arrayMap := make(map[int]int)
	arrayList := make([]int, 0)
	for _, num := range nums {
		_, ok := arrayMap[num]
		if !ok {
			arrayMap[num] = 1
			arrayList = append(arrayList, num)
		} else {
			arrayMap[num]++
		}
	}

	// // 快速排序sort.Slice
	// sort.Slice(arrayList, func(i, j int) bool {
	// 	return arrayMap[arrayList[i]] > arrayMap[arrayList[j]]
	// })

	// 快速排序myself
	var quicksort func(arrayMap map[int]int, datas []int, left, right int)
	quicksort = func(arrayMap map[int]int, datas []int, left, right int) {
		// 退出条件
		if left >= right {
			return
		}

		// 取最后一位当中间值
		middleNum := arrayMap[datas[right]]
		// start从最左侧的外边开始（left-1）
		start := left - 1
		// 从left开始遍历，遇到比middleNum小的值，start+1，然后datas[j]跟datas[start]交换位置
		for j := left; j < right; j++ {
			if arrayMap[datas[j]] > middleNum {
				start++
				datas[j], datas[start] = datas[start], datas[j]
			}
		}
		// 遍历完成后，先将start右侧的值跟middleNum的值交换，从而，start+1左侧的值比start+1小，start+1右侧的值比start的大，形成两个区域
		datas[start+1], datas[right] = datas[right], datas[start+1]
		// 两个区域继续执行，直到达到退出条件
		quicksort(arrayMap, datas, left, start)
		quicksort(arrayMap, datas, start+2, right)
	}
	quicksort(arrayMap, arrayList, 0, len(arrayList)-1)

	if len(nums) < 2 || k > len(nums) {
		return arrayList
	}

	return arrayList[:k]
}
