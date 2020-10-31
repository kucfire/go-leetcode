/*
	leetcode tag:242 title:有效字母异位词
*/

package leetcode242

// IsAnagram : method of count
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	arrayListS := make([]int, 26)
	arrayListT := make([]int, 26)
	for i := 0; i < len(s); i++ {
		arrayListS[s[i]-'a']++
		arrayListT[t[i]-'a']++
	}
	for i := 0; i < len(arrayListS); i++ {
		if arrayListS[i] != arrayListT[i] {
			return false
		}
	}
	return true
}
