/*
	leetcode tag:009 title:回文数
*/

package leetcode009

// normal method of myself
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tempA := x
	var reverse int
	for tempA > 0 {
		tempB := tempA % 10
		reverse = reverse*10 + tempB
		tempA = tempA / 10
	}
	return x == reverse
}
