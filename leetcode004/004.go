/*
	leetcode tag:004 title:寻找两个正序数组的中位数
*/
package leetcode004

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 合并数组
	numslen := len(nums1) + len(nums2)

	// 获取最小值的内置函数
	getmin := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

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

			// 当中间索引的值为1时，证明无法继续进行二分法
			if k == 1 {
				return float64(getmin(nums1[index1], nums2[index2]))
			}

			// 二分法处理
			half := k / 2
			// 如果索引超过原数组的长度，取数组的长度-1当做索引
			newIndex1 := getmin(index1+half, len(nums1)) - 1
			newIndex2 := getmin(index2+half, len(nums2)) - 1

			// 取出索引对应的数值进行处理
			pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
			if pivot1 <= pivot2 { // 当nums2对应的数值大于等于num1对应的数值,k值减少newindex与index的差值并加上1，因为k代表的是长度，必须要比index的大1
				k -= (newIndex1 - index1 + 1)
				index1 = newIndex1 + 1 //
			} else { // 当num1对应的数值大于num2对应的数值,k值
				k -= (newIndex2 - index2 + 1)
				index2 = newIndex2 + 1
			}
		}
		return float64(0)
	}

	// 长度判断
	if numslen%2 == 1 {
		midindex := numslen / 2
		return getKthElement(nums1, nums2, midindex+1)
	} else {
		midindex1, midindex2 := numslen/2, numslen/2+1
		return (getKthElement(nums1, nums2, midindex1) + getKthElement(nums1, nums2, midindex2)) / 2.0
	}
}
