package leetcode240

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
