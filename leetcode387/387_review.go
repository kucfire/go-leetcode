package leetcode387

func FirstUniqCharReview(s string) int {
	sArray := make([]int, 26)
	for i := 0; i < len(s); i++ {
		sArray[s[i]-'a'] = i
	}

	for i := 0; i < len(s); i++ {
		if i == sArray[s[i]-'a'] {
			return i
		} else {
			sArray[s[i]-'a'] = -1
		}
	}
	return -1
}
