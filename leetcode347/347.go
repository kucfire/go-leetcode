package leetcode347

func TopKFrequent(nums []int, k int) []int {
	// if len(nums) < 2 || k

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
	mv := func(datas []int, start, move int) int {

	}
	quicksort := func(datas []int, left, right int) {

	}
	quicksort(arrayList, 0, len(arrayList)-1)

	if len(nums) < 2 || k > len(nums) {
		return arrayList
	}

	return arrayList[:k]
}
