/*
	leetcode tag:344 title:反转字符串
*/

package leetcode344

// ReverseString : normal method
func ReverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
