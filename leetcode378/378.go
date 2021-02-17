/*
	leetcode tag:378 title:有序矩阵中第k小的元素
*/

package leetcode378

import "sort"

// method of myself
func KthSmallest(matrix [][]int, k int) int {
	arrayList := make([]int, 0)
	for _, son := range matrix {
		arrayList = append(arrayList, son...)
	}

	sort.Ints(arrayList)
	return arrayList[k-1]
}

// 二分查找
func KthSmallest2(matrix [][]int, k int) int {
	// 由于题干明确表明是n*n的矩阵，所以纵横都是n
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]

	// 这里面的n就是矩阵的纵横长度n
	check := func(matrix [][]int, mid, k, n int) bool {
		i, j := n-1, 0
		num := 0
		for i >= 0 && j < n {
			if matrix[i][j] <= mid {
				num += i + 1
				// 横向递增
				j++
			} else {
				// 纵向递减
				i--
			}
		}
		return num >= k
	}

	for left < right {
		mid := left + (right-left)/2
		if check(matrix, mid, k, n) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
