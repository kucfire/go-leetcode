/*
	leetcode tag:036 title:有效数独
*/

package leetcode036

// IsValidSudoku : normal method
func IsValidSudoku(board [][]byte) bool {
	var row, col, sbox [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				indexBox := (i/3)*3 + j/3
				if row[i][num] == 1 {
					return false
				} else {
					row[i][num] = 1
				}

				if col[j][num] == 1 {
					return false
				} else {
					col[j][num] = 1
				}
				if sbox[indexBox][num] == 1 {
					return false
				} else {
					sbox[indexBox][num] = 1
				}
			}
		}
	}
	return true
}
