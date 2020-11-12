/*
	leetcode tag:118 title:杨辉三角
*/

package leetcode118

// Generate : normal method
func Generate(numRows int) [][]int {
	var result [][]int
	for i := 0; i < numRows; i++ {
		result = append(result, []int{})
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				result[i] = append(result[i], 1)
			} else {
				result[i] = append(result[i], result[i-1][j-1]+result[i-1][j])
			}
		}
	}

	return result
}
