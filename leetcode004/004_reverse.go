package leetcode004

func ReverseFindMedianSortedArrays(nums1, nums2 []int) float64 {
	// 获取两个数组的总长度
	totalLen := len(nums1) + len(nums2)

	// 获取最小值的内置函数
	getMin := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	// 二分法获取中位数
	getKthElement := func(nums1, nums2 []int, k int) float64 {
		index1, index2 := 0, 0
		for {
			// 跳出条件
			if index1 == len(nums1) {
				return float64(nums2[index2+k-1])
			}
			if index2 == len(nums2) {
				return float64(nums1[index1+k-1])
			}

			if k == 1 {
				return float64(getMin(nums1[index1], nums2[index2]))
			}
			mid := k / 2

			newIndex1 := getMin(index1+mid, len(nums1)) - 1
			newIndex2 := getMin(index2+mid, len(nums2)) - 1
			pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
			if pivot1 <= pivot2 {
				k -= (newIndex1 - index1 + 1)
				index1 = newIndex1 + 1
			} else {
				k -= (newIndex2 - index2 + 1)
				index2 = newIndex2 + 1
			}

		}
		return 0.0
	}

	// 执行逻辑
	if totalLen%2 == 1 {
		midindex := totalLen / 2 // 下标
		return getKthElement(nums1, nums2, midindex+1)
	} else {
		midindex1, midindex2 := totalLen/2, totalLen/2+1
		return (getKthElement(nums1, nums2, midindex1) + getKthElement(nums1, nums2, midindex2)) / 2.0
	}
}
