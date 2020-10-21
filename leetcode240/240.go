package leetcode240

//normal
func SearchMatrix(matrix [][]int, target int) bool {
	for _, v := range matrix {
		for _, v0 := range v {
			if target == v0 {
				return true
			} else if target < v0 {
				continue
			}
		}
	}
	return false
}

//binary search
func SearchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	x, y, mid := 0, len(matrix[0])-1, 0
	for i := 0; i < len(matrix); i++ {
		for x <= y {
			mid = x + (y-x)/2
			if matrix[i][mid] == target {
				return true
			} else if matrix[i][mid] < target {
				x = mid + 1
			} else if matrix[i][mid] > target {
				y = mid - 1
			}
		}
		y = x - 1
		x = 0
	}
	return false
}

//
