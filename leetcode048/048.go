/*
	leetcode tag:048 title:旋转数组
*/

package leetcode048

// Rotate : method of normal
func Rotate(matrix [][]int) {
	size := len(matrix) //
	// size - 1 := size - 1
	for i := 0; i < size/2; i++ {
		for j := i; j < size-1-i; j++ {
			matrix[size-1-i][size-1-j], matrix[j][size-1-i] = matrix[j][size-1-i], matrix[size-1-i][size-1-j]
			matrix[i][j], matrix[j][size-1-i] = matrix[j][size-1-i], matrix[i][j]
			matrix[i][j], matrix[size-1-j][i] = matrix[size-1-j][i], matrix[i][j]
		}
	}
}
