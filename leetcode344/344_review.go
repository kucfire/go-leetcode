package leetcode344

func ReverseStringReview(s []byte) {
	n := len(s)
	for start, end := 0, n-1; start <= end; {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
