// /*
// 	leetcode tag:004 title:寻找两个正序数组的中位数
// */
package leetcode004

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)

	getMin := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	// 获取中位数元素
	getKthElement := func(nums1, nums2 []int, k int) int {
		index1, index2 := 0, 0
		for {
			// 跳出条件
			if index1 == len(nums1) {
				return nums2[index2+k-1]
			}
			if index2 == len(nums2) {
				return nums1[index1+k-1]
			}
			if k == 1 {
				return getMin(nums1[index1], nums2[index2])
			}

			//
			half := k / 2
			newIndex1 := getMin(index1+half, len(nums1)) - 1 // -1是因为里面的值受到长度的影响，而长度跟
			newIndex2 := getMin(index2+half, len(nums2)) - 1

			pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
			if pivot1 <= pivot2 {
				k -= (newIndex1 - index1 + 1)
				index1 = newIndex1 + 1
			} else {
				k -= (newIndex2 - index2 + 1)
				index2 = newIndex2 + 1
			}
		}
		return 0
	}

	// 分奇数长度和偶数长度的情况进行逻辑判断
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1)+getKthElement(nums1, nums2, midIndex2)) / 2.0
	}
	return 0
}
