package leetcode240

//
func SearchMatrixReview(matrix [][]int, target int) bool {
	// 二分查找
	// 第一层只能直接循环
	n := len(matrix)
	m := len(matrix[0])
	start, end, mid := 0, m-1, 0
	for i := 0; i < n; i++ {
		if target < matrix[i][0] {
			return false
		}
		for start <= end {
			mid = start + (end-start)/2
			if matrix[i][mid] > target {
				end = mid - 1
			} else if matrix[i][mid] < target {
				start = mid + 1
			} else {
				return true
			}
		}
		// 重置
		end = m - 1
		start = 0
	}
	return false
}
