package leetcode387

func FirstUniqCharReview(s string) int {
	sMap := make(map[rune]int)
	for _, str := range s {
		sMap[str]++
	}

	for i, str := range s {
		if sMap[str] == 1 {
			return i
		}
	}
	return -1
}
