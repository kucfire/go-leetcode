/*
	leetcode tag:005 title:最长回文子串
*/

package leetcode005

//method of double points
func longestPalindrome(s string) string {
	start, end := 0, 0
	// expend from single center or double
	expendArroundCenter := func(s string, left, right int) (int, int) {
		for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
		}
		return left + 1, right - 1
	}
	for i := 0; i <= len(s); i++ {
		left1, right1 := expendArroundCenter(s, i, i)
		left2, right2 := expendArroundCenter(s, i, i+1)
		if (right1 - left1) > (end - start) {
			start, end = left1, right1
		}
		if (right2 - left2) > (end - start) {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}
