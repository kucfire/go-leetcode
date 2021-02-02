package leetcode212

func FindWords(board [][]byte, words []string) []string {
	type arrayByte []byte
	result := make([]string, 0)
	m := len(board)    //1
	n := len(board[0]) //2

	// 将words转换成map，用来标识该word是否在board内出现过
	wordsMap := make(map[string]bool)
	for _, word := range words {
		wordsMap[word] = false
	}

	check := func(i, j int, word []byte) bool {
		signal := -1
		for index, b := range word {
			if index == 0 {
				// 第一个字母检测过了，可以跳过
				continue
			}
			if i-1 >= 0 && board[i-1][j] == b && signal != i-1 {
				signal = i
				i = i - 1
				continue
			} else if i+1 < m && board[i+1][j] == b && signal != i+1 {
				signal = i
				i = i + 1
				continue
			} else if j-1 >= 0 && board[i][j-1] == b && signal != j-1 {
				signal = j
				j = j - 1
				continue
			} else if j+1 < n && board[i][j+1] == b && signal != j+1 {
				signal = j
				j = j + 1
				continue
			}
			return false
		}
		return true
	}

	for word := range wordsMap {
		// find the first element
		firstElement := word[0]
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				// i = 0; j = 1
				if firstElement == board[i][j] && check(i, j, []byte(word)) {
					wordsMap[word] = true
				}
			}
		}
	}

	for key, value := range wordsMap {
		if value {
			result = append(result, key)
		}
	}

	return result
}
